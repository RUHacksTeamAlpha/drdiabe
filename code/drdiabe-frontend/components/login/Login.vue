<template>
  <div>
    <!-- <h1>{{ this.isAuthenticated() }}</h1> -->
    <v-window
      v-if="!isAuthenticated()"
      v-model="selectedItem"
      class="elevation-0"
    >
      <v-window-item>
        <v-card max-width="500" elevation="5" color="info" class="pa-4">
          <v-card-subtitle class="text-overline py-1">
            Welcome back
          </v-card-subtitle>
          <v-card-text
            class="text-md-h2 text-h4 py-0 font-weight-medium white--text"
          >
            Login
          </v-card-text>
          <v-card-subtitle class="text-body-2 text-md-body-1">
            Please enter your credentials below to login, or, press 'create an
            account' if you do not have one.
          </v-card-subtitle>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
                v-model="email"
                :rules="emailRules"
                label="E-mail"
                required
              ></v-text-field>

              <v-text-field
                v-model="password"
                :counter="10"
                :rules="passwordRules"
                label="Password"
                required
                type="password"
              ></v-text-field>

              <v-checkbox
                v-model="checkbox"
                label="Keep me signed in"
              ></v-checkbox>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn
              :disabled="!valid"
              :loading="loading"
              color="success"
              class="mr-4"
              @click="signIn"
            >
              Login
              <v-icon>mdi-login</v-icon>
            </v-btn>
            <v-btn color="primary" class="mr-4" @click="switchWindow">
              Create an account
            </v-btn>
            <v-btn color="error" class="mr-4" @click="reset"> Clear </v-btn>
          </v-card-actions>
        </v-card>
      </v-window-item>
      <v-window-item>
        <v-card max-width="500" elevation="5" color="info" class="pa-4">
          <v-card-subtitle class="text-overline py-1"> Welcome</v-card-subtitle>
          <v-card-text
            class="text-md-h2 text-h4 py-0 font-weight-medium white--text"
          >
            Create an Account
          </v-card-text>
          <v-card-subtitle class="text-body-2 text-md-body-1">
            Please fill out the form below and get started using drdiabe.tech
            today
          </v-card-subtitle>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
                v-model="name"
                :counter="10"
                :rules="nameRules"
                label="Name"
                required
              ></v-text-field>

              <v-text-field
                v-model="email"
                :rules="emailRules"
                label="E-mail"
                required
              ></v-text-field>

              <v-text-field
                v-model="password"
                :counter="10"
                :rules="passwordRules"
                label="Password"
                required
                type="password"
              ></v-text-field>

              <v-checkbox
                v-model="checkbox"
                :rules="[(v) => !!v || 'You must agree to continue!']"
                label="I agree to create an account"
                required
              ></v-checkbox>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn
              :disabled="!valid"
              :loading="loading"
              color="success"
              class="mr-4"
              @click="signUp"
            >
              Create
              <v-icon>mdi-pencil</v-icon>
            </v-btn>
            <v-btn color="primary" class="mr-4" @click="switchWindow">
              Switch to Login
            </v-btn>
            <v-btn color="error" class="mr-4" @click="reset"> Clear </v-btn>
          </v-card-actions>
        </v-card>
      </v-window-item>
    </v-window>
    <div v-if="isAuthenticated()">
      <v-btn @click="signOut()">Sign Out</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    loading: false,
    loader: null,
    selectedItem: 0,
    name: '',
    nameRules: [
      (v) => !!v || 'Name is required',
      (v) => (v && v.length <= 10) || 'Name must be less than 10 characters',
    ],
    password: '',
    passwordRules: [
      (v) => !!v || 'Password is required',
      (v) =>
        (v && v.length <= 20) || 'Password must be less than 20 characters',
    ],
    email: '',
    emailRules: [
      (v) => !!v || 'E-mail is required',
      (v) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    select: null,
    checkbox: false,
  }),
  watch: {
    loader() {
      const l = this.loader
      this[l] = !this[l]

      setTimeout(() => (this[l] = false), 3000)

      this.loader = null
    },
  },
  methods: {
    validate() {
      if (this.$refs.form.validate()) {
        this.loader = 'loading'
      }
    },
    reset() {
      this.$refs.form.reset()
    },
    switchWindow() {
      if (this.selectedItem === 0) {
        this.selectedItem = 1
      } else {
        this.selectedItem = 0
      }
      this.reset()
    },
    async signIn() {
      if (this.$refs.form.validate()) {
        try {
          await this.$fire.auth.signInWithEmailAndPassword(
            this.email,
            this.password
          )
          this.$fire.auth.currentUser.getIdToken().then((idToken) => {
            console.log(idToken)
          })
          this.$router.replace('/')
        } catch (error) {
          console.log(error)
        }
      }
    },
    async signUp() {
      if (this.$refs.form.validate()) {
        try {
          await this.$fire.auth.createUserWithEmailAndPassword(
            this.email,
            this.password
          )
          this.$fire.auth.currentUser.getIdToken().then((idToken) => {
            console.log(idToken)
          })
          this.$router.replace('/')
        } catch (e) {
          console.log(e)
        }
      }
    },
    async signOut() {
      try {
        await this.$fire.auth.signOut()
        console.log(this.isAuthenticated())
        console.log('Logged out')
        this.$router.replace('/')
      } catch (e) {
        alert(e)
      }
    },
    isAuthenticated() {
      return this.$fire.auth.currentUser !== null
    },
  },
  computed: {
    authUser() {
      return this.$fire.auth.currentUser
    },
  },
}
</script>

<style></style>
