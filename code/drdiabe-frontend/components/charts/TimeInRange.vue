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
    <v-divider></v-divider>
    <v-card-actions>
      <v-menu offset-y>
        <template v-slot:activator="{ on, attrs }">
          <v-btn color="primary" dark v-bind="attrs" v-on="on">
            Range
            <v-icon class="mx-1">mdi-arrow-expand-horizontal</v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-item v-for="(item, index) in items" :key="index">
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
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
      items: [{ title: 'Daily' }, { title: 'Weekly' }, { title: 'Monthly' }],
    }
  },
}
</script>
