<template>
  <v-container fluid>
    <v-row>
      <v-col cols="2" v-if="$vuetify.breakpoint.mdAndUp">
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
                  :disabled="selectedItem == 2"
                  color="accent"
                  dark
                  v-bind="attrs"
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
            <v-btn
              @click="download()"
              :disabled="selectedItem == 2"
              plain
              color="success"
            >
              DOWNLOAD
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-row class="my-2" v-if="$vuetify.breakpoint.smAndDown">
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
                    <v-btn color="accent" dark v-bind="attrs" v-on="on">
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
                <v-btn plain color="info"> DOWNLOAD </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
        <v-window v-model="selectedItem" class="elevation-0" vertical>
          <v-window-item>
            <!-- View Charts -->
            <v-row>
              <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
                <BGTotal
                  v-scroll-reveal="{ delay: 150 }"
                  :numberOfBG="numberOfBG"
                />
              </v-col>
              <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
                <TimeSince
                  v-scroll-reveal="{ delay: 550 }"
                  time="1 day, 2 hours, 15 minutes"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
                <Trends v-scroll-reveal="{ delay: 750 }" />
              </v-col>
              <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
                <TimeInRange v-scroll-reveal="{ delay: 350 }" />
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
      </v-col>
      <v-sheet
        color="transparent"
        elevation="0"
        :height="$vuetify.breakpoint.mdAndUp ? '400' : '100'"
        width="100%"
      >
      </v-sheet>
    </v-row>
  </v-container>
</template>

<script>
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
  data() {
    return {
      selectedItem: 0,
      items: [
        { text: 'View Charts', icon: 'mdi-chart-line' },
        { text: 'Summary', icon: 'mdi-calendar' },
        { text: 'Upload B.G', icon: 'mdi-water' },
      ],
      range: [{ title: 'Daily' }, { title: 'Weekly' }, { title: 'Monthly' }],
    }
  },
  computed: {
    numberOfBG() {
      return 23
    },
    TimeSince() {
      return '1 day, 2 hours, 15 minutes'
    },
  },
  methods: {
    download() {
      return null
    },
  },
}
</script>
