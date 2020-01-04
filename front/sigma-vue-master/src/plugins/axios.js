import Vue from 'vue'
import axios from 'axios'

// axios.defaults.baseURL = "https://curso-vue-50e2a.firebaseio.com/"

Vue.use({
  install(Vue) {
    // Vue.prototype.$http = axios
    Vue.prototype.$http = axios.create({
      baseURL: process.env.NODE_ENV === 'production' ? 'http://167.114.124.28:8081' : 'http://localhost:8081',
    
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