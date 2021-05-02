<template>
  <v-card>
    <v-stepper v-model="curr" alt-labels>
      <v-stepper-header>
        <v-stepper-step
          v-for="(step, n) in steps"
          :key="n"
          :color="stepStatus(n + 1)"
          :complete="stepComplete(n + 1)"
          :rules="[(value) => !!step.valid]"
          :step="n + 1"
        >
          {{ step.name }}
        </v-stepper-step>
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content step="1">
          <v-card class="mb-6 pa-2" height="200px">
            <v-card-text>
              Enter the Blood Glucose value you would like to upload. Please
              remember to use mmol/L as the unit for your B.G value
              <v-form :ref="'form1'" v-model="steps[0].valid" lazy-validation>
                <v-row class="my-2">
                  <v-icon color="error"> mdi-water </v-icon>
                  <v-text-field
                    v-model="bg"
                    :rules="steps[0].rules"
                    class="mx-4"
                    label="Blood Glucose (mmol/L)"
                  >
                  </v-text-field>
                </v-row>
              </v-form>
            </v-card-text>
          </v-card>
          <v-btn color="accent" @click="validate(0)"> Continue </v-btn>
        </v-stepper-content>

        <v-stepper-content step="2">
          <v-card class="mb-6 pa-2" height="200px">
            <v-card-text>
              Enter the date and time which corresponds with when this Blood
              Glucose value was taken.

              <v-form :ref="'form2'" v-model="steps[1].valid" lazy-validation>
                <v-dialog
                  ref="dialog1"
                  v-model="modal"
                  :return-value.sync="date"
                  persistent
                  width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="date"
                      label="Select a date"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :rules="steps[1].rules"
                    ></v-text-field>
                  </template>
                  <v-date-picker v-model="date" scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="accent" @click="modal = false">
                      Cancel
                    </v-btn>
                    <v-btn
                      text
                      color="accent"
                      @click="$refs.dialog1.save(date)"
                    >
                      OK
                    </v-btn>
                  </v-date-picker>
                </v-dialog>
                <v-dialog
                  ref="dialog"
                  v-model="modal2"
                  :return-value.sync="time"
                  persistent
                  width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="time"
                      label="Select a time"
                      prepend-icon="mdi-clock-time-four-outline"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :rules="steps[1].rules"
                    ></v-text-field>
                  </template>
                  <v-time-picker v-if="modal2" v-model="time" full-width>
                    <v-spacer></v-spacer>
                    <v-btn text color="accent" @click="modal2 = false">
                      Cancel
                    </v-btn>
                    <v-btn text color="accent" @click="$refs.dialog.save(time)">
                      OK
                    </v-btn>
                  </v-time-picker>
                </v-dialog>
              </v-form>
            </v-card-text>
          </v-card>
          <v-btn color="accent" @click="validate(1)"> Continue </v-btn>
          <v-btn text color="info" @click="reset()"> Cancel </v-btn>
        </v-stepper-content>

        <v-stepper-content step="3">
          <v-card class="mb-12">
            <v-card-text>
              Review your upload below and if everything is correct, press the
              upload button!
            </v-card-text>
            <v-row class="pa-2">
              <v-card-text class="text-overline white--text">
                Blood Glucose Value
                <v-chip color="accent">
                  <v-icon color="red">mdi-water</v-icon>{{ this.bg }}
                </v-chip>
              </v-card-text>
            </v-row>
            <v-row class="pa-2">
              <v-card-text class="text-overline white--text">
                Selected Date
                <v-chip color="accent">
                  <v-icon color="white">mdi-calendar</v-icon>{{ this.date }}
                </v-chip>
              </v-card-text>
            </v-row>
            <v-row class="pa-2">
              <v-card-text class="text-overline white--text">
                Selected Time
                <v-chip color="accent">
                  <v-icon color="white">mdi-clock</v-icon>{{ this.time }}
                </v-chip>
              </v-card-text>
            </v-row>
          </v-card>
          <v-btn color="success" @click="submit()"> Upload </v-btn>
          <v-btn text color="info" @click="reset()"> Cancel </v-btn>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      curr: 1,
      modal2: false,
      modal: false,
      steps: [
        {
          name: 'Enter B.G',
          rules: [
            (v) => (!!v && !isNaN(v) && v >= 0 && v <= 20.0) || 'Required.',
          ],
          valid: true,
        },
        {
          name: 'Date and Time',
          rules: [(v) => !!v || 'Required'],
          valid: true,
        },
        {
          name: 'Review',
          valid: true,
        },
      ],
      form1: [],
      form2: [],
      bg: null,
      valid: false,
      time: null,
      date: null,
    }
  },
  methods: {
    submit() {
      this.curr = 1
      console.log(this.date)
      console.log(this.time)
      console.log(this.date + 'T' + this.time + ':00Z')
    },
    stepComplete(step) {
      return this.curr > step
    },
    stepStatus(step) {
      return this.curr > step ? 'green' : 'blue'
    },
    validate(n) {
      this.steps[n].valid = false
      let v = null
      if (n === 0) {
        v = this.$refs.form1.validate()
      } else {
        v = this.$refs.form2.validate()
      }
      if (v) {
        this.steps[n].valid = true
        // continue to next
        this.curr += 1
      }
    },
    done() {
      this.curr = 4
    },
    reset() {
      this.$refs.form1.reset()
      this.$refs.form2.reset()
      this.curr = 1
    },
  },
}
</script>

<style></style>
