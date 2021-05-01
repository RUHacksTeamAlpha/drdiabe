<template>
  <v-container fluid>
    <v-row no-gutters>
      <v-col>
        <v-card class="ma-2" color="transparent" elevation="0">
          <v-card-text
            v-scroll-reveal="{ delay: 150 }"
            class="text-md-h2 text-h4 font-weight-medium white--text"
          >
            Test. Upload. Understand your diabetes.
          </v-card-text>

          <v-card-subtitle
            v-scroll-reveal="{ delay: 350 }"
            class="text-overline"
          >
            leveraging cloud computing to bring more affordable diabetic care to
            millions around the world
          </v-card-subtitle>

          <v-card-actions v-scroll-reveal="{ delay: 450 }">
            <v-btn outlined rounded color="accent"> Get started </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-card
          v-scroll-reveal="{ delay: 550 }"
          class="ma-2"
          color="transparent"
          elevation="0"
        >
          <v-img
            class="meter"
            contain
            height="400px"
            src="/img/meter.jpg"
          ></v-img>
        </v-card>
      </v-col>
    </v-row>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '400' : '150'"
      width="100%"
    >
    </v-sheet>
    <v-row>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
        <v-card
          v-scroll-reveal="{ delay: 150 }"
          class="rounded-card pa-0 up"
          color="transparent"
          elevation="0"
        >
          <v-img height="453px" src="/img/diverse-happy.jpg"></v-img>
        </v-card>
      </v-col>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
        <v-card v-scroll-reveal="{ delay: 450 }" class="pa-2 down">
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
              <v-spacer></v-spacer>
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
      :height="$vuetify.breakpoint.mdAndUp ? '400' : '100'"
      width="100%"
    >
    </v-sheet>
    <v-sheet color="primary" elevation="0" width="100%">
      <v-row>
        <v-col :cols="$vuetify.breakpoint.mdAndUp ? '6' : '12'">
          <v-card class="ma-2" color="transparent" elevation="0">
            <v-card-subtitle
              v-scroll-reveal="{ delay: 150 }"
              class="text-overline py-1"
            >
              Manage with drdiabe.tech
            </v-card-subtitle>
            <v-card-text
              v-scroll-reveal="{ delay: 150 }"
              class="text-md-h3 text-h4 font-weight-medium white--text"
            >
              Ready? Get started in just a few simple steps.
            </v-card-text>
            <v-card-subtitle
              v-scroll-reveal="{ delay: 150 }"
              class="text-subtitle-2 text-md-subtitle-1"
            >
              It's as easy as 1, 2, 3! Whether you want to use all of our
              features, or only some, we're here to help.
            </v-card-subtitle>
            <v-card-text>
              <v-row>
                <v-col cols="12">
                  <v-card
                    v-scroll-reveal="{ delay: 350 }"
                    color="secondary"
                    elevation="5"
                    class="pa-2"
                  >
                    <v-card-title class="pa-0">
                      1. Create an account
                    </v-card-title>
                    <v-card-text>
                      Start by creating an account which will act as place to
                      store your results
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12">
                  <v-card
                    v-scroll-reveal="{ delay: 550 }"
                    color="secondary"
                    elevation="5"
                    class="pa-2"
                  >
                    <v-card-title class="pa-0">
                      2. Upload some data
                    </v-card-title>
                    <v-card-text>
                      Upload your past blood glucose data or start fresh!
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12">
                  <v-card
                    v-scroll-reveal="{ delay: 750 }"
                    color="secondary"
                    elevation="5"
                    class="pa-2"
                  >
                    <v-card-title class="pa-0">
                      3. View your trends
                    </v-card-title>
                    <v-card-text>
                      Now you can see your daily / weekly / monthly trends with
                      our graphical interface!
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </v-card-text>
            <v-card-actions class="px-4" v-scroll-reveal="{ delay: 950 }">
              <v-btn outlined rounded color="accent"> Get started </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col v-if="$vuetify.breakpoint.mdAndUp" cols="6">
          <v-card v-scroll-reveal="{ delay: 150 }" class="mx-3 rounded-card">
            <v-img contain src="/img/testing.jpg"> </v-img>
          </v-card>
        </v-col>
      </v-row>
    </v-sheet>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '300' : '150'"
      width="100%"
    >
    </v-sheet>
  </v-container>
</template>

<script>
export default {}
</script>

<style scoped>
.rounded-card {
  border-radius: 50px;
}

.maxwidth {
  max-width: 100vw;
  padding: 0px;
}

@media (min-width: 800px) {
  .down {
    transform: translateY(100px);
  }

  .up {
    transform: translateY(-100px);
  }

  .meter {
    animation: fadeIn 1s;
  }
  @keyframes fadeIn {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
  @-webkit-keyframes fadeIn {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
}
</style>
