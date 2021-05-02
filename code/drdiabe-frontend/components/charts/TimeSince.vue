<template>
  <v-card class="pa-2" elevation="5">
    <v-row>
      <v-card-text class="text-overline text-center grey--text">
        <v-icon color="grey">mdi-clock</v-icon>
        Time Since Last Upload
        <div class="text-h5 white--text">{{ time }}</div>
      </v-card-text>
    </v-row>
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
