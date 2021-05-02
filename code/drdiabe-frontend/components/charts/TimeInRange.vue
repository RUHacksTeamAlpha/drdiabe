<template>
  <v-card class="pa-2" elevation="5">
    <v-card-text class="text-md-h4 text-h6 font-weight-medium white--text">
      {{ title }}
    </v-card-text>
    <v-card-subtitle class="py-0">
      {{ subtitle }}
    </v-card-subtitle>
    <v-card-text>
      <apexchart
        type="donut"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      title: 'Time in Range',
      subtitle: 'Check how often you are within your target B.G range',
      series: [12, 71, 17],
      chartOptions: {
        colors: ['#f0a30a', '#60AFFF', '#F44336'], // #60AFFF
        labels: ['Below Target', 'In Range', 'Above Target'],
        legend: {
          labels: {
            colors: 'white',
          },
        },
      },
    }
  },
}
</script>
