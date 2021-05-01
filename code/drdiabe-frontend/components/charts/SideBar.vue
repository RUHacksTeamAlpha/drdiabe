<template>
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
  </v-card>
</template>

<script>
export default {
  data: () => ({
    selectedItem: 0,
    items: [
      { text: 'View Charts', icon: 'mdi-chart-line' },
      { text: 'Summary', icon: 'mdi-file-chart' },
      { text: 'Upload B.G', icon: 'mdi-water' },
      { text: 'Review My Settings', icon: 'mdi-account' },
    ],
  }),
}
</script>

<style></style>
