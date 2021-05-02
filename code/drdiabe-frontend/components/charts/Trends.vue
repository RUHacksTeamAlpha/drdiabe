<template>
  <v-card class="pa-2" elevation="5">
    <v-card-text class="text-md-h4 text-h6 font-weight-medium white--text">
      {{ title }}
    </v-card-text>
    <v-card-subtitle class="py-0">
      {{ subtitle }}
    </v-card-subtitle>
    <v-card-text class="black--text">
      <apexchart type="line" :options="options" :series="series"></apexchart>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      title: 'Blood Glucose Trends',
      subtitle: 'Review your daily / weekly / monthly B.G trends',
      low: 4.0,
      high: 9.0,
      options: {
        colors: ['#60AFFF', '#f0a30a', '#F44336'], // #60AFFF
        stroke: {
          curve: 'smooth',
        },
        xaxis: {
          categories: [
            '0:00am',
            '3:00am',
            '6:00am',
            '9:00am',
            '12:00pm',
            '3:00pm',
            '6:00pm',
            '9:00pm',
          ],
        },
        yaxis: {
          max: 12,
          min: 2.0,
        },
      },
      series: [
        {
          name: 'My Blood Glucose',
          data: [6.4, 6.2, 7.2, 8.9, 5.3, 4.2, 6.0, 7.1],
        },
        {
          name: 'low',
          data: [4.0, 4.0, 4.0, 4.0, 4.0, 4.0, 4.0, 4.0],
        },
        {
          name: 'high',
          data: [9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0, 9.0],
        },
      ],
    }
  },
}
</script>
