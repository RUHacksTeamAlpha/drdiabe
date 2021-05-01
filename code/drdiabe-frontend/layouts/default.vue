<template>
  <v-app dark>
    <v-navigation-drawer v-model="drawer" temporary app>
      <v-list nav>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar fixed app>
      <v-app-bar-nav-icon
        v-if="$vuetify.breakpoint.smAndDown"
        @click.stop="drawer = !drawer"
      />
      <div>
        <v-toolbar-title>
          <router-link to="/" tag="span" style="cursor: pointer">
            <v-icon>mdi-stethoscope</v-icon>
            {{ $vuetify.breakpoint.mdAndUp ? title : trunc }}
          </router-link>
        </v-toolbar-title>
      </div>
      <v-spacer />
      <v-tabs v-if="$vuetify.breakpoint.mdAndUp" color="white">
        <v-spacer></v-spacer>
        <v-tabs-slider color="white"></v-tabs-slider>
        <v-tab v-for="item in items" :key="item.id" :to="item.to">
          <v-icon class="ma-2" color="grey darken-2">{{ item.icon }}</v-icon>
          {{ item.title }}
          <v-icon class="ma-2" color="transparent">{{ item.icon }}</v-icon>
          <!-- dummy icon to center the text instead of the space between icon and text -->
        </v-tab>
      </v-tabs>
    </v-app-bar>
    <v-main>
      <v-container>
        <v-fade-transition mode="out-in" :hide-on-leave="true">
          <nuxt />
        </v-fade-transition>
      </v-container>
    </v-main>
    <v-footer absolute dark padless app color="primary">
      <v-card flat tile width="100%" class="white--text text-center">
        <v-card-text>
          <v-icon> mdi-stethoscope </v-icon>
        </v-card-text>
        <v-card-text class="pa-0 text-body-2 text-md-body-1 white--text">
          drdiabe.tech
        </v-card-text>
        <v-card-text>
          <v-tooltip v-for="(item, index) in social" :key="index" bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                :href="item.href"
                target="_blank"
                icon
                class="mx-4 white--text"
              >
                <v-icon size="24px" color="white" v-bind="attrs" v-on="on">
                  {{ item.icon }}
                </v-icon>
              </v-btn>
            </template>
            <span>{{ item.alt }}</span>
          </v-tooltip>
        </v-card-text>
      </v-card>
      <v-divider></v-divider>
      <v-card-text>
        <v-row>
          <v-col align="center" justify="center">
            &copy; {{ new Date().getFullYear() }}
          </v-col>
        </v-row>
      </v-card-text>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Home',
          to: '/',
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'My Hub',
          to: '/hub',
        },
        {
          icon: 'mdi-lock',
          title: 'Login/Logout',
          to: '/login',
        },
        {
          icon: 'mdi-account',
          title: 'Account',
          to: '/account',
        },
      ],
      title: 'drdiabe.tech',
      trunc: 'dr.d',
      social: [
        {
          icon: 'mdi-facebook',
          alt: 'Facebook',
          href: '',
        },
        {
          icon: 'mdi-instagram',
          alt: 'Instagram',
          href: '',
        },
        {
          icon: 'mdi-github',
          alt: 'Github',
          href: '',
        },
      ],
    }
  },
  computed: {
    getItems() {
      return this.$store.state.routes.getItems
    },
  },
  methods: {},
}
</script>
