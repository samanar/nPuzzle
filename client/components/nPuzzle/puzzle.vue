<template>
  <div>
    <v-card class="puzzle_container" :class="{'puzzle_container_win': win}">
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
    export default {
        name: "puzzle",
        props: ['rowSize', 'colSize'],
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
                this.numbers = this.shuffleArray(this.numbers);
            },
            shuffleArray(arr) {
                let newArr = arr.slice();
                for (let i = newArr.length - 1; i > 0; i--) {
                    let j = Math.floor(Math.random() * (i + 1));
                    let temp = newArr[i];
                    newArr[i] = newArr[j];
                    newArr[j] = temp;
                }
                return newArr;
            },
            generateNumbers() {
                this.numbers = [];
                let max = this.rowSize * this.colSize;
                for (let i = 0; i < max; i++) this.numbers.push(i);
                this.shuffle();
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
                        this.$emit('addLog', 'empty cell moved up');
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i - 1, coordinates.j);
                        break;
                    case 'down':
                        this.$emit('addLog', 'empty cell moved down');
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i + 1, coordinates.j);
                        break;
                    case 'left':
                        this.$emit('addLog', 'empty cell moved left');
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i, coordinates.j - 1);
                        break;
                    case 'right' :
                        this.$emit('addLog', 'empty cell moved right');
                        this.swapTwoNumbers(coordinates.i, coordinates.j, coordinates.i, coordinates.j + 1);
                        break;
                }
            },
            swapTwoNumbers(first_i, first_j, second_i, second_j) {
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
        },
        watch: {
            win(newVal) {
                this.$emit('setWin', newVal);
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
    max-width: 500px;
    display: flex;
    flex-wrap: wrap;
  }

  .puzzle_container_win {
    border: 5px solid green;
  }

  .black_box {
    background-color: black;
    border: none;
    border-radius: 0 !important;
  }

  .black_box_win {
    color: green;
    background-color: green;
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

</style>
