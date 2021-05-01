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
      <v-tabs color="white">
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
        <nuxt />
      </v-container>
    </v-main>
    <v-footer absolute dark padless app>
      <v-card flat tile width="100%" class="white--text text-center">
        <v-card-text class="pa-1">
          &copy; {{ title }} - {{ new Date().getFullYear() }}
        </v-card-text>
      </v-card>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: false,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Welcome',
          to: '/',
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Inspire',
          to: '/inspire',
        },
      ],
      miniVariant: false,
      title: 'drdiabe.tech',
      trunc: 'dr.d',
    }
  },
  computed: {
    getItems() {
      return this.$store.state.routes.getItems
    },
  },
}
</script>
