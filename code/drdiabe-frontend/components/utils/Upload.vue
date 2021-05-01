<template>
  <v-card>
    <v-stepper v-model="e1" alt-labels>
      <v-stepper-header>
        <v-stepper-step :complete="e1 > 1" step="1">
          Enter B.G Value
        </v-stepper-step>

        <v-divider></v-divider>

        <v-stepper-step :complete="e1 > 2" step="2">
          Enter Time
        </v-stepper-step>

        <v-divider></v-divider>

        <v-stepper-step step="3"> Confirm </v-stepper-step>
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content step="1">
          <v-card class="mb-6 pa-2" height="200px">
            <v-card-text>
              Enter the Blood Glucose value you would like to upload. Please
              remember to use mmol/L as the unit for your B.G value
            </v-card-text>
            <v-text-field
              v-model="bg"
              class="mx-4"
              label="Blood Glucose (mmol/L)"
            >
            </v-text-field>
          </v-card>

          <v-btn color="primary" @click="e1 = 2"> Continue </v-btn>

          <v-btn text color="error"> Cancel </v-btn>
        </v-stepper-content>

        <v-stepper-content step="2">
          <v-card class="mb-6 pa-2" height="200px">
            <v-card-text>
              Enter the date and time which corresponds with when this Blood
              Glucose value was taken.
            </v-card-text>
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
                ></v-text-field>
              </template>
              <v-date-picker v-model="date" scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="modal = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.dialog1.save(date)">
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
                ></v-text-field>
              </template>
              <v-time-picker v-if="modal2" v-model="time" full-width>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="modal2 = false">
                  Cancel
                </v-btn>
                <v-btn text color="primary" @click="$refs.dialog.save(time)">
                  OK
                </v-btn>
              </v-time-picker>
            </v-dialog>
          </v-card>

          <v-btn color="primary" @click="e1 = 3"> Continue </v-btn>
          <v-btn text color="error"> Cancel </v-btn>
        </v-stepper-content>

        <v-stepper-content step="3">
          <v-card class="mb-12" height="250px">
            <v-card-text>
              Review your upload below and if everything is correct, press the
              upload button!
            </v-card-text>
            <v-card-text class="text-overline white--text">
              Blood Glucose Value: {{ bg }}
            </v-card-text>
            <v-card-text class="text-overline white--text">
              Selected Time: {{ time }}
            </v-card-text>
            <v-card-text class="text-overline white--text">
              Selected Date: {{ date }}
            </v-card-text>
          </v-card>

          <v-btn color="success" @click="e1 = 1"> Upload </v-btn>

          <v-btn text color="error"> Cancel </v-btn>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      e1: 1,
      bg: null,
      time: null,
      modal2: false,
      modal: false,
      date: new Date().toISOString().substr(0, 10),
    }
  },
}
</script>

<style></style>