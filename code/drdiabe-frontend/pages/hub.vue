<template>
  <v-container fluid>
    <div v-if="$fire.auth.currentUser !== null">
      <v-row class="my-3">
        <v-col cols="12">
          <v-card class="pa-2" elevation="5">
            <v-subheader class="text-overline">NAVIGATION</v-subheader>
            <v-divider></v-divider>
            <v-list>
              <v-list-item-group v-model="selectedItem" color="primary">
                <v-list-item v-for="(item, i) in items" :key="i">
                  <v-list-item-icon>
                    <v-icon v-text="item.icon"></v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title v-text="item.text"></v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list-item-group>
            </v-list>
            <v-divider></v-divider>
            <v-card-actions>
              <v-menu offset-y>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="accent"
                    dark
                    v-bind="attrs"
                    :disabled="selectedItem === 2"
                    v-on="on"
                  >
                    Range
                    <v-icon class="mx-1">mdi-arrow-expand-horizontal</v-icon>
                  </v-btn>
                </template>
                <v-list>
                  <v-list-item v-for="(item, index) in range" :key="index">
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
              <v-spacer />
              <v-btn plain color="success" :disabled="selectedItem === 2">
                DOWNLOAD
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
      <v-window v-model="selectedItem" class="elevation-0" vertical>
        <v-window-item>
          <!-- View Charts -->
          <v-row>
            <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
              <BGTotal :numberOfBG="data ? data.length : 0" />
            </v-col>
            <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
              <TimeSince :time="data[0] ? data[data.length - 1].time : '0'" />
            </v-col>
          </v-row>
          <v-row>
            <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
              <Trends />
            </v-col>
            <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
              <TimeInRange />
            </v-col>
          </v-row>
        </v-window-item>
        <v-window-item>
          <!-- Summary -->
          <v-row>
            <v-col> <Summary /> </v-col>
          </v-row>
        </v-window-item>
        <v-window-item>
          <!-- Upload B.G -->
          <v-row>
            <v-col>
              <Upload />
            </v-col>
          </v-row>
        </v-window-item>
      </v-window>
      <v-sheet
        color="transparent"
        elevation="0"
        :height="$vuetify.breakpoint.mdAndUp ? '200' : '100'"
        width="100%"
      >
      </v-sheet>
    </div>
    <div v-else>
      <v-row justify="center">
        <v-card v-scroll-reveal class="pa-2">
          <v-card-text class="text-h4 text-md-h2 white--text">
            User not logged in
          </v-card-text>
          <v-card-subtitle> Please login to view this page </v-card-subtitle>
          <v-card-actions>
            <v-btn outlined color="primary" to="/login"> Login </v-btn>
          </v-card-actions>
        </v-card>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import axios from 'axios'
import Trends from '../components/charts/Trends'
import TimeInRange from '../components/charts/TimeInRange'
import BGTotal from '../components/charts/BGTotal'
import TimeSince from '../components/charts/TimeSince'
import Upload from '../components/utils/Upload'
import Summary from '../components/utils/Summary'

// import SideBar from '../components/charts/SideBar'

export default {
  components: {
    Trends,
    TimeInRange,
    BGTotal,
    TimeSince,
    // SideBar,
    Upload,
    Summary,
  },
  mounted() {
    if (this.$fire.auth.currentUser !== null) {
      this.getData()
    }
  },
  data() {
    return {
      selectedItem: 0,
      items: [
        { text: 'View Charts', icon: 'mdi-chart-line' },
        { text: 'Summary', icon: 'mdi-calendar' },
        { text: 'Upload B.G', icon: 'mdi-water' },
      ],
      data: [],
      range: [{ title: 'Daily' }, { title: 'Weekly' }, { title: 'Monthly' }],
    }
  },
  methods: {
    download() {
      return null
    },
    async getData() {
      const domain = 'https://ruhacks-backend-c6jokpb2dq-uc.a.run.app/'
      const token = await this.$fire.auth.currentUser.getIdToken()
      await axios
        .get(domain + '/data', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          this.data = res.data
          console.log(res.data)
        })
        .catch((err) => {
          console.log('unable to load data: ' + err)
        })
    },
    getTimeSince() {
      return null
    },
  },
}
</script>
