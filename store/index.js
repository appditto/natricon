export const state = () => ({
  clientID: 0
})

export const mutations = {
  SET_ID(state, id) {
    let intId = parseInt(id)
    if (isNaN(intId)) {
        intId = 0
    }
    state.clientID = intId
  }
}

export const actions = {
  SET_ID({ commit }, id) {
    commit('SET_ID', id)
  }
}