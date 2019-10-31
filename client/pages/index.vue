<template>
  <v-row justify="center">
    <v-col cols="12" sm="12" md="6" class="px-6">
      <puzzle></puzzle>
    </v-col>
    <v-col cols="12" sm="12" md="6" class="px-6">
      hello
    </v-col>
  </v-row>
</template>

<script>

    import puzzle from "../components/nPuzzle/puzzle";

    export default {
        name: 'indexPage',
        components: {puzzle},
    }
</script>
