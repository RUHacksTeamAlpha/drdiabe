<template>
  <v-card class="pa-2" elevation="5">
    <v-card-text class="text-overline text-center white--text py-1">
      Time of Last Upload
    </v-card-text>
    <v-card-text class="text-center py-1">
      <v-chip rounded outlined color="info" class="text-h5">
        <v-icon color="white" class="mx-1">mdi-clock</v-icon>{{ time }}</v-chip
      >
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    time: String,
  },
}
</script>

<style></style>
