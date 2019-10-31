<template>
  <div>
    <v-row justify="center">
      <v-col cols="12" sm="12" md="6" class="px-6">
        <puzzle ref="puzzle"
                :rowSize="rowSize"
                :colSize="colSize"
                v-on:incrementMoves="incrementMoves"
                v-on:addLog="addLog"
                v-on:setWin="setWin"
        ></puzzle>
      </v-col>
      <v-col cols="12" sm="12" md="6" class="px-6">
        <puzzle-controller
          :win="win"
          :moves="moves"
          v-on:changeLevel="changeLevel"
          v-on:newPuzzleGame="newPuzzleGame">
        </puzzle-controller>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" sm="12">
        <v-textarea
          outlined
          flat
          rows="5"
          height="150px"
          no-resize
          readonly
          ref="logs"
          filled
          label="logs"
          class="logs"
          :value="logs"
        ></v-textarea>
      </v-col>
    </v-row>
  </div>
</template>

<script>
    import puzzle from "../components/nPuzzle/puzzle";
    import puzzleController from "../components/nPuzzle/puzzleController";

    export default {
        name: 'indexPage',
        components: {puzzle, puzzleController},
        data() {
            return {
                win: false,
                moves: 0,
                rowSize: 3,
                colSize: 3,
                logs: ''
            }
        },
        methods: {
            incrementMoves() {
                this.moves++;
            },
            addLog(text) {
                this.logs += text + '\n';
                this.$refs.logs.scrollTop = this.$refs.logs.scrollHeight;
            },
            changeLevel(level) {
                this.rowSize = level;
                this.colSize = level;
                this.moves = 0;
            },
            newPuzzleGame() {
                console.log("here");
                this.moves = 0;
                this.$refs.puzzle.shuffle();
            },
            setWin(win) {
                this.win = win
            }
        }
    }
</script>

<style>
  textarea {
    color: #7f828b !important;
  }
</style>
