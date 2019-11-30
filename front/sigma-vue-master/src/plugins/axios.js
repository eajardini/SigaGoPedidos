import Vue from 'vue'
import axios from 'axios'

// axios.defaults.baseURL = "https://curso-vue-50e2a.firebaseio.com/"

Vue.use({
  install(Vue) {
    // Vue.prototype.$http = axios
    Vue.prototype.$http = axios.create({
      baseURL: "http://localhost:8081",
      // headers: {
      //     "Authorization":"meu token"
      // }s
    })

    Vue.prototype.$borodin = axios.create({
      baseURL: "https://curso-vue-50e2a.firebaseio.com/"
    })

    Vue.prototype.$http.interceptors.request.use(config => {
      // console.log(config.method)
      return config
    }, error => Promise.reject(error))

    Vue.prototype.$http.interceptors.response.use(resp => {
      // const array = []
      // for (let chave in resp.data) {
      //   array.push ({id: chave, ...resp.data[chave]})
      // }
      // resp.data = array
      return resp
    }, error => Promise.reject(error))

  }
})