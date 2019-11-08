<template>
  <v-card class="pa-5">
    <v-row>
      <v-col cols="12">
        <v-select
          outlined
          label="difficulty"
          v-model="selectedLevel"
          :items="levels"
          @change="changeLevel"
          :disabled="solving || showingDemo">
        </v-select>
        <v-select
          outlined
          label="solving algorithm"
          v-model="algorithm"
          :items="algorithms"
          @change="changeAlgorithm"
          :disabled="solving || showingDemo"></v-select>
        <v-btn
          color="primary"
          block
          large
          @click="emitNewPuzzleGame"
          :disabled="solving || showingDemo">
          Start new game
        </v-btn>
        <v-btn
          color="secondary"
          block
          large
          @click="solve"
          :loading="solving"
          class="mt-2"
          :disabled="showingDemo">
          Solve
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" sm="12" md="12">
        number of moves:
        <span class="value" :class="{'value_win': win}">{{moves}}</span>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
    export default {
        props: ['win', 'moves', 'solving', 'showingDemo'],
        name: "puzzleController",
        data() {
            return {
                selectedLevel: 3,
                algorithm: 'bfs',
                algorithms: [
                    {text: 'BFS', value: 'bfs'},
                    {text: 'DFS', value: 'dfs'},
                    {text: 'GREEDY', value: 'greedy'},
                    {text: 'A*', value: 'aStar'}
                ],
                levels: [
                    {text: 'easy', value: 3},
                    {text: 'medium', value: 4},
                    {text: 'hard', value: 5}
                ]
            }
        },
        methods: {
            emitNewPuzzleGame() {
                this.$emit('newPuzzleGame');
            },
            changeLevel() {
                console.log(this.selectedLevel);
                this.$emit('changeLevel', this.selectedLevel);
            },
            solve() {
                this.$emit('solve');
            },
            changeAlgorithm() {
                this.$emit('changeAlgorithm', this.algorithm);
            }
        }
    }
</script>

<style scoped>
  .value {
    text-decoration: underline;
    font-size: 1.1em;
    margin-left: 10px;
  }

  .value_win {
    color: green
  }
</style>
