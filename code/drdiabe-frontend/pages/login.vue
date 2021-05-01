<template>
  <v-container fluid>
    <v-row justify="center">
      <v-col cols="4">
        <Login v-scroll-reveal />
      </v-col>
    </v-row>
    <v-sheet
      color="transparent"
      elevation="0"
      :height="$vuetify.breakpoint.mdAndUp ? '100' : '50'"
      width="100%"
    >
    </v-sheet>
  </v-container>
</template>

<script>
import Login from '../components/login/Login'

export default {
  components: {
    Login,
  },
}
</script>
