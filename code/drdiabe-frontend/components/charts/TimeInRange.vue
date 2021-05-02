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
import axios from 'axios'

export default {
  data() {
    return {
      low: 4.5,
      high: 9.0,
      title: 'Time in Range',
      subtitle: 'Check how often you are within your target B.G range',
      series: [],
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
  mounted() {
    if (this.$fire.auth.currentUser !== null) {
      this.getTargets()
      this.getData()
    }
  },
  methods: {
    async getData() {
      const domain = 'https://ruhacks-backend-c6jokpb2dq-uc.a.run.app'
      const token = await this.$fire.auth.currentUser.getIdToken()
      await axios
        .get(domain + '/data', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          let i = 0
          const len = res.data.length
          let low = 0
          let high = 0
          let tir = 0
          for (; i < len; i++) {
            const bg = res.data[i].bg
            if (bg < this.low) {
              low += bg
            }
            if (bg >= this.low && bg <= this.high) {
              tir += bg
            }
            if (bg > this.high) {
              high += bg
            }
          }
          low /= len
          high /= len
          tir /= len
          this.series = []
          this.series.push(low)
          this.series.push(tir)
          this.series.push(high)
        })
        .catch((err) => {
          console.log('error getting chart data: ' + err)
        })
    },
    async getTargets() {
      const domain = 'https://ruhacks-backend-c6jokpb2dq-uc.a.run.app'
      const token = await this.$fire.auth.currentUser.getIdToken()
      await axios
        .get(domain + '/profile', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          this.low = res.data.low
          this.high = res.data.high
        })
        .catch((err) => {
          console.log('error getting high/low preference: ' + err)
        })
    },
  },
}
</script>
