import Vue from 'vue';
import App from './App.vue';
import router from './router';
import AutoComplete from 'primevue/autocomplete';
import Accordion from 'primevue/accordion';
import AccordionTab from 'primevue/accordiontab';
import Button from 'primevue/button';
import Calendar from 'primevue/calendar';
import Card from 'primevue/card';
import Chart from 'primevue/chart';
import Checkbox from 'primevue/checkbox';
import Chips from 'primevue/chips';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import DataView from 'primevue/dataview';
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions';
import Dialog from 'primevue/dialog';
import Dropdown from 'primevue/dropdown';
import Editor from 'primevue/editor';
import Fieldset from 'primevue/fieldset';
import FileUpload from 'primevue/fileupload';
import FullCalendar from 'primevue/fullcalendar';
import InputMask from 'primevue/inputmask';
import InputSwitch from 'primevue/inputswitch';
import InputText from 'primevue/inputtext';
import Listbox from 'primevue/listbox';
import Menu from 'primevue/menu';
import Message from 'primevue/message';
import MultiSelect from 'primevue/multiselect';
import OrderList from 'primevue/orderlist';
import OrganizationChart from 'primevue/organizationchart';
import OverlayPanel from 'primevue/overlaypanel';
import Paginator from 'primevue/paginator';
import Panel from 'primevue/panel';
import Password from 'primevue/password';
import PickList from 'primevue/picklist';
import ProgressBar from 'primevue/progressbar';
import Rating from 'primevue/rating';
import RadioButton from 'primevue/radiobutton';
import SelectButton from 'primevue/selectbutton';
import Slider from 'primevue/slider';
import Sidebar from 'primevue/sidebar';
import SplitButton from 'primevue/splitbutton';
import Spinner from 'primevue/spinner';
import Textarea from 'primevue/textarea';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';
import Toolbar from 'primevue/toolbar';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import ToggleButton from 'primevue/togglebutton';
import Tree from 'primevue/tree';
import TriStateCheckbox from 'primevue/tristatecheckbox';
import ValidationMessage from 'primevue/validationmessage';
import money from 'v-money';

import 'primevue/resources/themes/nova-light/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeflex/primeflex.css';
import 'primeicons/primeicons.css';
import '@fullcalendar/core/main.min.css';
import '@fullcalendar/daygrid/main.min.css';
import '@fullcalendar/timegrid/main.min.css';
import './assets/layout/layout.scss';



import "./plugins/axios"

import AppSubmenu from './AppSubmenu';
Vue.directive("appsubmenu", AppSubmenu)

Vue.use(ToastService);
Vue.use(money, {precision: 2})


Vue.config.productionTip = false;

Vue.component('Accordion', Accordion);
Vue.component('AccordionTab', AccordionTab);
Vue.component('AutoComplete', AutoComplete);
Vue.component('Button', Button);
Vue.component('Calendar', Calendar);
Vue.component('Card', Card);
Vue.component('Chart', Chart);
Vue.component('Checkbox', Checkbox);
Vue.component('Chips', Chips);
Vue.component('Column', Column);
Vue.component('DataTable', DataTable);
Vue.component('DataView', DataView);
Vue.component('DataViewLayoutOptions', DataViewLayoutOptions);
Vue.component('Dialog', Dialog);
Vue.component('Dropdown', Dropdown);
Vue.component('Editor', Editor);
Vue.component('Fieldset', Fieldset);
Vue.component('FileUpload', FileUpload);
Vue.component('FullCalendar', FullCalendar);
Vue.component('InputMask', InputMask)
Vue.component('InputSwitch', InputSwitch);
Vue.component('InputText', InputText);
Vue.component('Listbox', Listbox);
Vue.component('Menu', Menu);
Vue.component('Message', Message);
Vue.component('MultiSelect', MultiSelect);
Vue.component('OrderList', OrderList);
Vue.component('OrganizationChart', OrganizationChart);
Vue.component('OverlayPanel', OverlayPanel);
Vue.component('Paginator', Paginator);
Vue.component('Panel', Panel);
Vue.component('Password', Password);
Vue.component('ProgressBar', ProgressBar);
Vue.component('PickList', PickList);
Vue.component('RadioButton', RadioButton);
Vue.component('Rating', Rating);
Vue.component('SelectButton', SelectButton);
Vue.component('Slider', Slider);
Vue.component('Sidebar', Sidebar);
Vue.component('Spinner', Spinner);
Vue.component('SplitButton', SplitButton);
Vue.component('TabView', TabView);
Vue.component('TabPanel', TabPanel);
Vue.component('Textarea', Textarea);
Vue.component('Toast', Toast);
Vue.component('Toolbar', Toolbar);
Vue.component('ToggleButton', ToggleButton);
Vue.component('Tree', Tree);
Vue.component('TriStateCheckbox', TriStateCheckbox);
Vue.component('ValidationMessage', ValidationMessage);


new Vue({
	router,
	render: h => h(App)
}).$mount('#app');
