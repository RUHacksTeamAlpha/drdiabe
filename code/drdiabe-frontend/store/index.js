// store/index.js
export const state = () => ({
  authUser: {},
})
export const actions = {
  async onAuthStateChangedAction({ commit }, { authUser, claims }) {
    const { uid, email, emailVerified, displayName, photoURL } = authUser
    await commit('SET_USER', {
      uid,
      email,
      emailVerified,
      displayName,
      photoURL,
      isAdmin: claims.custom_claim,
    })
  },
  async nuxtServerInit({ dispatch, commit }, { res }) {
    console.log(res.locals)
    if (res && res.locals && res.locals.user) {
      const { allClaims: claims, idToken: token, ...authUser } = res.locals.user
      await dispatch('onAuthStateChangedAction', {
        authUser,
        claims,
        token,
      })
      commit('ON_AUTH_STATE_CHANGED_MUTATION', { authUser, claims, token })
    }
  },
}
export const mutations = {
  ON_AUTH_STATE_CHANGED_MUTATION(state, { authUser, claims }) {
    const { uid, email, emailVerified, displayName, photoURL } = authUser
    state.authUser = {
      uid,
      displayName,
      email,
      emailVerified,
      photoURL: photoURL || null,
      isAdmin: claims.custom_claim,
    }
  },
  SET_USER(state, payload) {
    console.log(payload)
    state.authUser = payload
  },
}
// end
