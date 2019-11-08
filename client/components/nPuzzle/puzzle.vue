<template>
  <div>
    <v-card class="puzzle_container" :class="{'puzzle_container_win': win}" style="position: relative">
      <v-card class="loading_container" v-show="solving">
        <robot></robot>
      </v-card>
      <v-card
        outlined
        v-for="(number ,index) in numbers"
        :key="number"
        @click="swapWithZero(index)"
        class="number_container puzzle_item"
        :style="{flex : '1 0 ' + puzzle_item_size + '%'}"
        :class="{'black_box' : number === 0, 'black_box_win': number === 0 && win , 'puzzle_item_win' : win}"
        :disabled="number === 0 || win"
      >
        <v-responsive :aspect-ratio="1">
          <v-card-title class="mx-auto my-auto text-center card_number">
            {{number}}
          </v-card-title>
        </v-responsive>
      </v-card>
    </v-card>
  </div>
</template>

<script>
    import robot from "~/components/general/robot";

    export default {
        name: "puzzle",
        props: ['rowSize', 'colSize', 'solving', 'showingDemo'],
        components: {robot},
        data() {
            return {
                numbers: [],
                win: false,
                puzzle_item_size: '33'
            }
        },
        created() {
            this.generateNumbers();
        },
        methods: {
            checkWin() {
                let mismatch = false;
                for (let i = 0; i < this.numbers.length - 1; i++) {
                    if (this.numbers[i] !== i + 1) {
                        mismatch = true;
                    }
                }
                this.win = !mismatch;
            },
            shuffle() {
                let numberOfMoves = this.getRandomInt(250);
                numberOfMoves += 50;
                for (let i = 0; i < numberOfMoves; i++) {
                    let randomMove = this.getRandomInt(5);
                    switch (randomMove) {
                        case 1:
                            this.moveZero('up');
                            break;
                        case 2:
                            this.moveZero('down');
                            break;
                        case 3:
                            this.moveZero('left');
                            break;
                        case 4:
                            this.moveZero('right');
                            break;
                    }
                }
            },
            getRandomInt(max) {
                let random = Math.floor(Math.random() * Math.floor(max));
                if (random === 0) random = 1;
                return random
            },
            generateNumbers() {
                this.numbers = [];
                let max = this.rowSize * this.colSize;
                for (let i = 1; i < max; i++) this.numbers.push(i);
                this.numbers.push(0);
                this.shuffle();
                this.$emit('resetCounter');
                switch (this.rowSize) {
                    case 3:
                        this.puzzle_item_size = 33;
                        break;
                    case 4:
                        this.puzzle_item_size = 25;
                        break;
                    case 5:
                        this.puzzle_item_size = 20;
                }
            },
            findZero() {
                for (let i = 0; i < this.rowSize; i++) {
                    for (let j = 0; j < this.colSize; j++) {
                        if (this.numbers[i * this.rowSize + j] === 0) {
                            return {
                                i: i,
                                j: j
                            }
                        }
                    }
                }
                return null;
            },
            convertIndexToCoordinates(index) {
                let i = Math.trunc(index / this.rowSize);
                let j = index - (i * this.rowSize);
                if (j < 0)
                    j = j * -1;
                return {
                    i: i,
                    j: j
                }
            },
            moveZero(direction) {
                let coordinates = this.findZero();
                if (!coordinates) {
                    console.log("coordinates couldn't be found");
                    return;
                }
                switch (direction) {
                    case 'up' :
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i - 1, coordinates.j);
                        break;
                    case 'down':
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i + 1, coordinates.j);
                        break;
                    case 'left':
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i, coordinates.j - 1);
                        break;
                    case 'right' :
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i, coordinates.j + 1);
                        break;
                }
            },
            swapTwoNumbers(first_i, first_j, second_i, second_j) {
                if (first_i >= this.rowSize || first_j >= this.colSize || first_i < 0 | first_j < 0) return;
                if (second_i >= this.rowSize || second_j >= this.colSize || second_i < 0 | second_j < 0) return;
                if (!this.isZeroCoordinates(first_i, first_j)) {
                    return;
                }
                let first_index = first_i * this.rowSize + first_j;
                let second_index = second_i * this.rowSize + second_j;
                let first_num = this.numbers[first_index];
                let second_num = this.numbers[second_index];
                this.$set(this.numbers, first_index, second_num);
                this.$set(this.numbers, second_index, first_num);
                this.$emit('incrementMoves');
            },
            isZeroCoordinates(i, j) {
                if (i >= this.rowSize || j >= this.colSize || i < 0 | j < 0) return false;
                return this.numbers[i * this.rowSize + j] === 0
            },
            swapWithZero(index) {
                let coordinates = this.convertIndexToCoordinates(index);
                let i = coordinates.i;
                let j = coordinates.j;
                if (this.isZeroCoordinates(i + 1, j)) this.moveZero('up');
                if (this.isZeroCoordinates(i - 1, j)) this.moveZero('down');
                if (this.isZeroCoordinates(i, j + 1)) this.moveZero('left');
                if (this.isZeroCoordinates(i, j - 1)) this.moveZero('right');
            },
            demoAnswer(answer) {
                this.$emit('setDemo', true);
                let self = this;
                (function show(i, length) {
                    setTimeout(function () {
                        self.$emit('incrementMoves');
                        let demo = answer[i].split('-').map(Number);
                        for (let i = 0; i < demo.length; i++) {
                            self.$set(self.numbers, i, demo[i]);
                        }
                        i++;
                        if (i < length) {
                            show(i, answer.length);
                        }
                    }, 500);
                })(0, answer.length);
            }
        },
        watch: {
            win(newVal) {
                this.$emit('setWin', newVal);
                this.$emit('setDemo', false);
            },
            numbers() {
                this.checkWin();
            },
            rowSize() {
                this.generateNumbers();
            }
        }
    }
</script>

<style scoped>
  .puzzle_item {
    display: flex;
    transition: transform 1s;
  }

  .puzzle_item_win {
    border: 2px solid green !important;
  }

  .puzzle_container {
    display: flex;
    flex-wrap: wrap;
    max-width: 600px;
    margin: 0 auto;
  }

  .puzzle_container_win {
    border: 5px solid green !important;
  }

  .black_box {
    background-color: black !important;
    border: none !important;
    border-radius: 0 !important;
  }

  .black_box_win {
    color: green !important;
    background-color: green !important;
  }


  .card_number {
    width: 100%;
    height: 100%;
    align-items: center;
    align-content: center;
    text-align: center;
    justify-content: center;
    font-size: 2rem;
  }

  .loading_container {
    z-index: 10;
    height: 100%;
    width: 100%;
    position: absolute;
    background-color: rgba(255, 255, 255, 0.5);
    top: 0;
    bottom: 0
  }

</style>
