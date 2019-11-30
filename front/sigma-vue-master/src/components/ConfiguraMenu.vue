<template>
	<div class="p-grid">
		<div class="p-col-12">
			<div class="card">
				<h1>Escolha qual menu você deseja alterar</h1>

				<div class="p-grid p-fluid">
					<div class="p-col-6">
						<div class="p-col-6">
							<Button type="button" label="Somente Financeiro e Componentes" @click="somenteFinanceiroECompoentes" />
							
						</div>
						<div class="p-col-6">
							<Button type="button" label="Menu Completo" @click="toggleDataTable" />
						
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	methods:{
		somenteFinanceiroECompoentes(){
			this.$menu =  [
				{
          "label": "Financeiro",          
          "items": [
            {
              "label": "Clientes",              
              "items": [
                {
                  "label": "Cadastro de Clientes",                  
                }
              ]
            }
          ]
        },


			]
		}
	}

}
</script>

<style scoped>

</style>