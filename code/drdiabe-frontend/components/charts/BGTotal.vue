<template>
  <v-card class="pa-2" elevation="5">
    <v-row>
      <v-card-text class="text-overline text-center grey--text">
        <v-icon color="red">mdi-water</v-icon>
        Total Number of B.G Records
        <div class="text-h5 white--text">{{ numberOfBG }}</div>
      </v-card-text>
    </v-row>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      numberOfBG: '28',
    }
  },
}
</script>

<style></style>
