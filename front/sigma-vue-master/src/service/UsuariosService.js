// import axios from 'axios'

export default class UsuariosService {

	getTodosUsuarios() {
	//	return axios.get('').then(res => res.data.data);
		return this.$http.get("/listaTodosUsuarios").then(res => res.data.resposta);
	}

}