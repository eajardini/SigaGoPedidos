import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from './components/Dashboard.vue';
import CadastroDeFuncionarios from './components/recursoshumanos/CadastroDeFuncionarios.vue';
import AtualizacaoDeFuncionarios from './components/recursoshumanos/AtualizacaoDeFuncionarios.vue';
import RelatorioDeFuncionarios from './components/recursoshumanos/RelatorioDeFuncionario.vue';
// import ConfiguraMenu from './components/ConfiguraMenu';



Vue.use(Router);

export default new Router({
	mode: "history",
	routes: [		
		{
			path: '/',
			name: 'dashboard',
			component: Dashboard
		},
		//Administração
		{
			path: '/manutencaodeusuario',
			name: 'manutencaodeusuario',
			component: () => import('./components/administração/ManutencaoDeUsuarios.vue') 
		},
		//Recursos Humanos
		{
			path: '/manutencaodefuncionario',
			name: 'manutencaodefuncionario',
			component: () => import('./components/recursoshumanos/ManutencaoDeFuncionarios.vue') 		
		},
		{
			path: '/cadastrodefuncionario',
			name: 'cadastrodefuncionario',
			component: CadastroDeFuncionarios
			// component: () => import('./components/recursoshumanos/CadastroDeFuncionarios.vue') 		
		},
		{
			path: '/atualizacaodefuncionario',
			name: 'atualizacaodefuncionario',
			component: AtualizacaoDeFuncionarios
			// component: () => import('./components/recursoshumanos/CadastroDeFuncionarios.vue') 		
		},
		{
			path: '/relatoriofuncionario',
			name: 'Relatorio de Funcionario',
			component: RelatorioDeFuncionarios
			// component: () => import('./components/recursoshumanos/CadastroDeFuncionarios.vue') 		
		},

		{
			path: '/sample',
			name: 'sample',
			component: () => import('./components/SampleDemo.vue')
		},
		{
			path: '/forms',
			name: 'forms',
			component: () => import('./components/FormsDemo.vue')
		},
		{
			path: '/data',
			name: 'data',
			component: () => import('./components/DataDemo.vue')
		},
		{
			path: '/overlays',
			name: 'overlays',
			component: () => import('./components/OverlaysDemo.vue')
		},
		{
			path: '/menus',
			name: 'menus',
			component: () => import('./components/MenusDemo.vue')
		},
		{
			path: '/messages',
			name: 'messages',
			component: () => import('./components/MessagesDemo.vue')
		},
		{
			path: '/charts',
			name: 'charts',
			component: () => import('./components/ChartsDemo.vue')
		},
		{
			path: '/misc',
			name: 'misc',
			component: () => import('./components/MiscDemo.vue')
		},
		{
			path: '/empty',
			name: 'empty',
			component: () => import('./components/EmptyPage.vue')
		},
		{
			path: '/documentation',
			name: 'documentation',
			component: () => import('./components/Documentation.vue')
		},
	],
	scrollBehavior() {
		return {x: 0, y: 0};
	}
});