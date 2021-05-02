<template>
  <v-container fluid>
    <div v-if="$fire.auth.currentUser !== null" class="pa-4">
      <v-row justify="center">
        <v-card color="transparent" elevation="0">
          <v-card-text class="text-h4 text-md-h2 white--text">
            My Account
          </v-card-text>
        </v-card>
      </v-row>
      <v-row justify="center">
        <v-col>
          <v-card class="mx-auto pa-0" max-width="374" elevation="5">
            <v-img
              v-if="$fire.auth.currentUser.photoURL !== null"
              :src="$fire.auth.currentUser.photoURL"
              contain
              tile
            >
            </v-img>
            <v-card-text class="text-h5 text-md-h4 white--text">
              {{ this.user.name }}
            </v-card-text>
            <v-divider></v-divider>
            <v-card-subtitle> Target Preferences </v-card-subtitle>
            <v-card-text class="text-overline">
              High
              <v-chip color="info darken-1">
                <v-icon color="red">mdi-water</v-icon>{{ this.user.low }}
              </v-chip>
            </v-card-text>
            <v-card-text class="text-overline">
              Low
              <v-chip color="info darken-1">
                <v-icon color="red">mdi-water</v-icon>{{ this.user.high }}
              </v-chip>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <v-row justify="center">
        <Profile />
      </v-row>
      <v-row justify="center">
        <v-btn class="mt-4" @click="printUser()">Print User Details</v-btn>
      </v-row>
    </div>
    <div v-else>
      <v-row justify="center">
        <v-card v-scroll-reveal class="pa-2">
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
import axios from 'axios'
import Profile from '../components/utils/Profile'

export default {
  components: {
    Profile,
  },
  async mounted() {
    if (this.$fire.auth.currentUser !== null) {
      const domain = 'https://ruhacks-backend-c6jokpb2dq-uc.a.run.app'
      const token = await this.$fire.auth.currentUser.getIdToken()
      await axios
        .get(domain + '/profile', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          this.user = res.data
          console.log(res.data)
        })
        .catch((err) => {
          console.log('error getting profile: ' + err)
        })
    }
  },
  data() {
    return {
      user: {
        name: '',
        age: 0,
        weight: 0,
        low: 0.0,
        high: 0.0,
      },
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
