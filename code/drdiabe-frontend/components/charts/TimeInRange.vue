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
          type="donut"
          :options="chartOptions"
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
      title: 'Range Timings',
      subtitle: 'Something goes here',
      series: [13, 71, 14],
      chartOptions: {
        labels: ['Below Target', 'In Range', 'Above Target'],
      },
    }
  },
}
</script>
