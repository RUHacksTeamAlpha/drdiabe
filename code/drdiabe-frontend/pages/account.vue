<template>
  <v-container fluid>
    <v-row>
      <v-card v-scroll-reveal color="transparent" elevation="0">
        <v-card-text class="text-h4 text-md-h2 white--text">
          My Account
        </v-card-text>
      </v-card>
    </v-row>
    <v-row justify="center">
      <v-col>
        <v-img></v-img>
      </v-col>
      <v-col cols="8">
        <v-card
          v-scroll-reveal="{ delay: 150 }"
          class="mx-auto pa-0"
          max-width="374"
          elevation="5"
        >
          <v-img contain :src="img" tile> </v-img>
          <v-card-text class="text-h5 text-md-h4 white--text">
            {{ name }}
          </v-card-text>
          <v-card-subtitle class="text-h6 py-0 white--text">
            Target Preferences
          </v-card-subtitle>
          <v-card-text>
            High: {{ high }} <br />
            Low: {{ low }}
          </v-card-text>
          <v-divider class="mx-2"></v-divider>
          <v-card-actions>
            <v-btn outlined color="primary"> Edit </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-img></v-img>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-btn @click="printUser()">Print User Details</v-btn>
    </v-row>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '100' : '50  '"
      width="100%"
    >
    </v-sheet>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      name: 'Taylor Schwarts',
      img: '/img/person.png',
      high: '9.8',
      low: '4.5',
    }
  },
  methods: {
    printUser() {
      console.log(this.$fire.auth.currentUser)
    },
  },
}
</script>

<style></style>
