<template>
  <v-container fluid>
    <div v-if="this.$fire.auth.currentUser !== null" class="pa-4">
      <v-row justify="center">
        <v-card v-scroll-reveal color="transparent" elevation="0">
          <v-card-text class="text-h4 text-md-h2 white--text">
            My Account
          </v-card-text>
        </v-card>
      </v-row>
      <v-row justify="center">
        <v-col>
          <v-card
            v-scroll-reveal="{ delay: 150 }"
            class="mx-auto pa-0"
            max-width="374"
            elevation="5"
          >
            <v-img
              contain
              v-if="this.$fire.auth.currentUser.photoURL !== null"
              :src="this.$fire.auth.currentUser.photoURL"
              tile
            >
            </v-img>
            <v-card-text class="text-h5 text-md-h4 white--text">
              {{ this.$fire.auth.currentUser.displayName }}
            </v-card-text>
            <v-divider></v-divider>
            <v-card-subtitle> Target Preferences </v-card-subtitle>
            <v-card-text class="text-overline">
              High
              <v-chip color="info darken-1">
                <v-icon color="red">mdi-water</v-icon>{{ low }}
              </v-chip>
            </v-card-text>
            <v-card-text class="text-overline">
              Low
              <v-chip color="info darken-1">
                <v-icon color="red">mdi-water</v-icon>{{ high }}
              </v-chip>
            </v-card-text>
            <v-divider class="mx-2"></v-divider>
            <v-card-actions>
              <v-btn outlined color="primary"> Edit </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
      <v-row justify="center">
        <Profile />
      </v-row>
      <v-row justify="center">
        <v-btn @click="printUser()" class="mt-4">Print User Details</v-btn>
      </v-row>
    </div>
    <div v-else>
      <v-row justify="center">
        <v-card v-scroll-reveal>
          <v-card-text class="text-h4 text-md-h2 white--text">
            User not logged in
          </v-card-text>
          <v-card-subtitle> Please login to view this page </v-card-subtitle>
          <v-card-actions>
            <v-btn outlined color="primary" to="/login"> Login </v-btn>
          </v-card-actions>
        </v-card>
      </v-row>
    </div>
    <v-sheet color="transparent" elevation="0" height="100" width="100%">
    </v-sheet>
  </v-container>
</template>

<script>
import Profile from '../components/utils/Profile'

export default {
  components: {
    Profile,
  },
  data() {
    return {
      high: '9.8',
      low: '4.5',
    }
  },
  methods: {
    printUser() {
      console.log(this.$fire.auth.currentUser)
    },
  },
  computed: {},
}
</script>

<style></style>
