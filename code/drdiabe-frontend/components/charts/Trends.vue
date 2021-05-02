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
import axios from 'axios'

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
          categories: [],
          labels: {
            style: {
              colors: [],
            },
          },
        },
        yaxis: {
          max: 15,
          min: 2.0,
          labels: {
            style: {
              colors: [],
            },
          },
        },
        legend: {
          labels: {
            colors: 'white',
          },
        },
      },
      series: [
        {
          name: 'My Blood Glucose',
          data: [],
        },
        {
          name: 'low',
          data: [],
        },
        {
          name: 'high',
          data: [],
        },
      ],
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
          for (; i < len; i++) {
            this.options.xaxis.categories.push(res.data[i].time) // not working
            this.options.xaxis.labels.style.colors.push('white')
            this.options.yaxis.labels.style.colors.push('white')
            this.series[0].data.push(res.data[i].bg)
            this.series[1].data.push(this.low) // should be user set low
            this.series[2].data.push(this.high) // should be user set high
          }
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
