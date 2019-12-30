<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Cadastro de Funcionarios</h1>
      </div>
      <div class="p-grid">
        <div class="p-col-10">
          <div class="card">
            <p>Nome:</p>

            </div>
        </div>
          <div align="right">
            <div class="p-col-8" >
              <div class="card">
                <h4></h4>
                <img witdh="100px" height="100px" v-bind:src="fotoFuncionario" />
                <br />
                <input @change="onFileChanged" type="file" id="selectedFile" style="display: none;" />
                <input
                  type="button"
                  value="Escolha uma foto"
                  onclick="document.getElementById('selectedFile').click();"
                />
              </div>
            </div>
          </div>
        
      </div>
    </div>
  </div>
</template>

<script>
//import UsuariosService from "../../service/UsuariosService";
export default {
  data() {
    return {
      filters: {},
      selectedFile: null,
      //files:null,
      fotoFuncionario: "",

      funcionarios: "",
      selectedFuncionarios1: null
    };
  },
  funcionariosService: null,
  methods: {
    onFileChanged(event) {
      this.selectedFile = event.target.files[0];
      this.onUpload();
    },
    onUpload() {
      //https://www.youtube.com/watch?v=J84uiTfjbGA
      const formData = new FormData();
      formData.append("foto", this.selectedFile, this.selectedFile.name);
      this.$http
        .post("/rh/upLoadFotoFuncionario", formData, { responseType: "blob" })
        .then(res => {
          let reader = new FileReader();
          reader.readAsDataURL(res.data);
          reader.onload = () => {
            this.fotoFuncionario = reader.result;
          };
        });

      // verificar porque a foto nao esta apareceno

      // this.$http.get("/rh/listaTodosFuncionarios").then(res => {
      //               this.funcionarios = res.data.resposta});
    },
    previewFiles() {
      this.files = this.$refs.myFiles.files;
      console.log("Valor de files", this.files);
      console.log("Valor de  refs files", this.$refs.myFiles.files);
    }
  },
  created() {
    // this.usuariosService = new UsuariosService();
  },
  mounted() {
    // this.usuariosService.getTodosUsuarios().then(data => (this.usuarios = data));
  }
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>