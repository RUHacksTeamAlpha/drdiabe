<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-card max-width="250">
          <v-row>
            <v-icon>mdi-water</v-icon>

            <v-card-text class="pa-0 mx-2">
              Total Number of B.G Records
            </v-card-text>

            <v-card-text class="pa-0 mx-2"> 7 </v-card-text>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-spacer></v-spacer>
      <v-col>
        <Trends />
      </v-col>
      <v-col>
        <TimeInRange />
      </v-col>
      <v-spacer></v-spacer>
    </v-row>
  </v-container>
</template>

<script>
import Trends from '../components/charts/Trends'
import TimeInRange from '../components/charts/TimeInRange'

export default {
  components: {
    Trends,
    TimeInRange,
  },
}
</script>
