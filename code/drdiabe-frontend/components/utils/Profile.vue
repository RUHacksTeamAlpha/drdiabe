<template>
  <v-card class="pa-2">
    <v-card-text class="text-md-h3 text-h5 white--text">
      Edit Profile
    </v-card-text>
    <v-card-text>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field
          v-model="name"
          :rules="nameRules"
          label="Name"
          required
        ></v-text-field>

        <v-text-field
          v-model="age"
          :rules="int"
          label="Age"
          required
        ></v-text-field>

        <v-text-field
          v-model="weight"
          :rules="int"
          label="Weight"
          required
        ></v-text-field>

        <v-text-field
          v-model="low"
          :rules="number"
          label="Low"
          required
        ></v-text-field>

        <v-text-field
          v-model="high"
          :rules="number"
          label="High"
          required
        ></v-text-field>

        <v-btn
          :disabled="!valid"
          outlined
          color="success"
          class="mr-4"
          @click="validate"
        >
          Update Profile
        </v-btn>

        <v-btn color="error" outlined class="mr-4" @click="reset">
          Reset Form
        </v-btn>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    name: '',
    age: 0,
    weight: 0,
    low: 0.0,
    high: 0.0,
    nameRules: [(v) => !!v || 'Name is required'],
    int: [(v) => Number.isInteger(v)], // !isNaN(v)
    number: [(v) => Number.isFinite(v)], // !isNaN(v)
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
    reset() {
      this.$refs.form.reset()
    },
  },
}
</script>

<style></style>
