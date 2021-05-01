export const state = () => ({
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
})

export const getters = () => ({
  getItems() {
    return this.items
  },
})
