<template>
  <div>
    <v-alert
      prominent
      type="error"
      v-if="error"
    >
      <v-row align="center">
        <v-col class="grow">{{error}}</v-col>
        <v-col class="shrink">
          <v-btn outlined @click="reconnect">
            <v-icon>mdi-refresh</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-alert>
    <v-row justify="center">
      <v-col cols="12" sm="12" md="6" class="px-6">
        <puzzle ref="puzzle"
                :rowSize="rowSize"
                :colSize="colSize"
                :solving="solving"
                :showingDemo="showingDemo"
                v-on:incrementMoves="incrementMoves"
                v-on:addLog="addLog"
                v-on:resetCounter="resetCounter"
                v-on:setWin="setWin"
                v-on:setDemo="setDemo"
        ></puzzle>
      </v-col>
      <v-col cols="12" sm="12" md="6" class="px-6">
        <puzzle-controller
          :win="win"
          :moves="moves"
          :solving="solving"
          :showingDemo="showingDemo"
          v-on:changeLevel="changeLevel"
          v-on:solve="solve"
          v-on:newPuzzleGame="newPuzzleGame">
        </puzzle-controller>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" sm="12">
        <v-card>
          <v-card-title>
            <span class="animated_dots">Logs</span>
          </v-card-title>
          <v-divider></v-divider>
          <div class="log_container" ref="log">
            <div v-for="(m, index) in logs" :key="index" class="log">
              {{index}} : {{m}}
            </div>
          </div>

        </v-card>
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
                socket: null,
                logs: [],
                answer: [],
                error: '',
                solving: false,
                showingDemo: false
            }
        },
        mounted() {
            let self = this;
            this.socket = new WebSocket('ws://localhost:8080/ws');

            this.socket.onopen = () => {
                this.error = '';
                self.socket.send(JSON.stringify({
                    title: 'hello',
                    type: 'nPuzzle',
                    body: '',
                    numbers: [1, 2, 3, 4, 5, 6, 7, 8, 9]
                }));
            };

            this.socket.onclose = (event) => {
                this.error = 'socket is closed !! click to try reconnecting'
            };

            this.socket.onerror = (err) => {
                console.error('socket error ', err);
            };

            this.socket.onmessage = (msg) => {
                let message = JSON.parse(msg.data);
                switch (message.title) {
                    case 'log' :
                        this.addLog('number of nodes expanded ' + message.body);
                        break;
                    case 'goal' :
                        this.solving = false;
                        this.$refs.puzzle.demoAnswer(message.solution);
                        break;
                    case 'test':
                        break;
                }
            };

        },
        methods: {
            setDemo(demo) {
                this.showingDemo = demo
            },
            reconnect() {
                window.location.reload()
            },
            incrementMoves() {
                this.moves++;
            },
            addLog(text) {
                this.logs.push(text)
            },
            resetCounter() {
                this.moves = 0;
            },
            solve() {
                this.solving = true;
                this.socket.send(JSON.stringify({
                    type: 'nPuzzle',
                    title: 'init',
                    rowSize: this.rowSize,
                    colSize: this.colSize,
                    numbers: this.$refs.puzzle.numbers
                }));
            },
            changeLevel(level) {
                this.rowSize = level;
                this.colSize = level;
                this.moves = 0;
            },
            newPuzzleGame() {
                this.moves = 0;
                this.$refs.puzzle.generateNumbers();
            },
            setWin(win) {
                this.win = win
            },
        },
        watch: {
            logs: function () {
                this.$refs.log.scrollTop = this.$refs.log.scrollHeight
            }
        }
    }
</script>

<style>
  textarea {
    color: #7f828b !important;
  }

  .log_container {
    height: 200px !important;
    max-height: 200px !important;
    overflow-y: scroll;
  }

  .log {
    padding-left: 20px;
    color: #47494E;
    font-size: 1.1em;
  }

  .animated_dots:after {
    content: ' .';
    animation: dots 1s steps(5, end) infinite;
  }

  @keyframes dots {
    0%, 20% {
      color: rgba(0, 0, 0, 0);
      text-shadow: .25em 0 0 rgba(0, 0, 0, 0),
      .5em 0 0 rgba(0, 0, 0, 0);
    }
    40% {
      color: #000000;
      text-shadow: .25em 0 0 rgba(0, 0, 0, 0),
      .5em 0 0 rgba(0, 0, 0, 0);
    }
    60% {
      text-shadow: .25em 0 0 #000000,
      .5em 0 0 rgba(0, 0, 0, 0);
    }
    80%, 100% {
      text-shadow: .25em 0 0 #000000,
      .5em 0 0 #000000;
    }
  }

</style>
