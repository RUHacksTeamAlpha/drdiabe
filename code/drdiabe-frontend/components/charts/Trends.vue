<template>
  <v-card class="pa-2" elevation="5">
    <v-card-text class="text-md-h4 text-h6 font-weight-medium white--text">
      {{ title }}
    </v-card-text>
    <v-card-subtitle class="py-0">
      {{ subtitle }}
    </v-card-subtitle>
    <v-card-text>
      <v-container fluid>
        <apexchart
          width="500"
          type="line"
          :options="options"
          :series="series"
        ></apexchart>
      </v-container>
    </v-card-text>
    <v-divider></v-divider>
    <v-card-actions>
      <v-btn rounded outlined color="accent"> UPLOAD B.G </v-btn>
      <v-spacer />
      <v-btn plain color="info"> DOWNLOAD </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      title: 'My Blood Glucose Trends',
      subtitle: 'Something goes here',
      options: {
        chart: {
          id: 'vuechart-example',
        },
        xaxis: {
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998],
        },
      },
      series: [
        {
          name: 'series-1',
          data: [30, 40, 45, 50, 49, 60, 70, 91],
        },
      ],
    }
  },
}
</script>
