<template>
  <v-card class="pa-2" elevation="5">
    <v-card-text class="text-overline text-center white--text py-1">
      Total Number of B.G Records
    </v-card-text>
    <v-card-text class="text-center py-1">
      <v-chip rounded outlined color="info" class="text-h5">
        <v-icon color="red">mdi-water</v-icon>{{ numberOfBG }}</v-chip
      >
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    numberOfBG: Number,
  },
  methods: {
    async getTotalNumber() {},
  },
}
</script>

<style></style>
