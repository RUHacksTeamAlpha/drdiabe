<template>
  <v-container fluid>
    <v-row no-gutters>
      <v-col>
        <v-card class="ma-2" color="transparent" elevation="0">
          <v-slide-y-reverse-transition appear>
            <v-card-text
              class="text-md-h2 text-h4 font-weight-medium white--text"
            >
              Test. Upload. Understand your diabetes.
            </v-card-text>
          </v-slide-y-reverse-transition>
          <v-fade-transition appear>
            <v-card-subtitle class="text-overline">
              leveraging cloud computing to bring more affordable diabetic care
              to millions around the world
            </v-card-subtitle>
          </v-fade-transition>
          <v-card-actions>
            <v-btn outlined rounded color="accent"> Get started </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '300' : '150'"
      width="100%"
    >
    </v-sheet>
    <v-row>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
        <v-card class="rounded-card pa-0 up" color="transparent" elevation="0">
          <v-img height="453px" src="/img/diverse-happy.jpg"></v-img>
        </v-card>
      </v-col>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
        <v-card class="pa-2 down">
          <v-card-title class="text-md-h3 text-h4 font-weight-medium">
            Features
          </v-card-title>
          <v-card-subtitle class="text-md-subtitle-1 text-subtitle-2">
            Maybe you're a good Diabetic. Maybe you have a handwritten logbook
            and you're wondering how drdiabe.tech can help you. Don't worry,
            we've got you covered. Here's a list of key feature you get access
            to when you sign up for an account with drdiabe.tech!
          </v-card-subtitle>
          <v-divider></v-divider>
          <v-card-text>
            <v-card elevation="0" class="pa-2">
              <v-card-title class="pa-0 text-md-h5 text-h6">
                <v-icon class="mx-2">mdi-upload</v-icon>
                Upload your Blood Glucose
              </v-card-title>
              <v-card-text class="pa-0"> Enter your B.G and ... </v-card-text>
            </v-card>
            <v-card elevation="0" class="pa-2">
              <v-card-title class="pa-0">
                <v-icon class="mx-2">mdi-chart-bell-curve</v-icon>
                View your Blood Glucose Trends <br />
              </v-card-title>
              <v-card-text class="pa-0">
                Monitor your Blood Glucose Trends ...
              </v-card-text>
            </v-card>
            <v-card elevation="0" class="pa-2">
              <v-card-title class="pa-0">
                <v-icon class="mx-2">mdi-clipboard-outline</v-icon>
                Generate Summary Reports <br />
              </v-card-title>
              <v-card-text class="pa-0">
                Generate a report which details the average trends of your blood
                glucose ...
              </v-card-text>
            </v-card>
          </v-card-text>
          <v-card-actions>
            <v-row>
              <v-col>
                <v-btn outlined rounded color="accent"> Get started </v-btn>
              </v-col>
              <v-col>
                <v-btn rounded plain color="info">
                  See how drdiabe.tech was made
                </v-btn>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-spacer v-if="$vuetify.breakpoint.mdAndUp"></v-spacer>
    </v-row>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '300' : '150'"
      width="100%"
    >
    </v-sheet>
    <v-row> </v-row>
  </v-container>
</template>

<script>
export default {}
</script>

<style scoped>
.rounded-card {
  border-radius: 50px;
}

@media (min-width: 800px) {
  .down {
    transform: translateY(15px);
  }

  .up {
    transform: translateY(-15px);
  }
}
</style>
