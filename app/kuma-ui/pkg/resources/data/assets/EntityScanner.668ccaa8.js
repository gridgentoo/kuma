import{C as g,o as s,k as n,l as o,t as w,y as r,z as m,H as c,j as f,F as v,n as S,m as b,cU as k,a as h,w as y,b as C}from"./index.563e1198.js";const z={name:"FormFragment",props:{title:{type:String,required:!1,default:null},forAttr:{type:String,required:!1,default:null},allInline:{type:Boolean,default:!1},hideLabelCol:{type:Boolean,default:!1},equalCols:{type:Boolean,default:!1},shiftRight:{type:Boolean,default:!1}}},I={class:"form-line-wrapper"},E={key:0,class:"form-line__col"},q=["for"];function B(e,d,t,u,a,_){return s(),n("div",I,[o("div",{class:m(["form-line",{"has-equal-cols":t.equalCols}])},[t.hideLabelCol?r("",!0):(s(),n("div",E,[o("label",{for:t.forAttr,class:"k-input-label"},w(t.title)+": ",9,q)])),o("div",{class:m(["form-line__col",{"is-inline":t.allInline,"is-shifted-right":t.shiftRight}])},[c(e.$slots,"default")],2)],2)])}const se=g(z,[["render",B],["__scopeId","data-v-03ab9d16"]]);const $={props:{steps:{type:Array,default:()=>{}},sidebarContent:{type:Array,required:!0,default:()=>{}},footerEnabled:{type:Boolean,default:!0},nextDisabled:{type:Boolean,default:!0}},emits:["goToStep"],data(){return{start:0}},computed:{step:{get(){return this.steps[this.start].slug},set(e){return this.steps[e].slug}},indexCanAdvance(){return this.start>=this.steps.length-1},indexCanReverse(){return this.start<=0}},mounted(){this.setStartingStep()},methods:{goToNextStep(){this.start++,this.updateQuery("step",this.start),this.$emit("goToStep",this.step)},goToPrevStep(){this.start--,this.updateQuery("step",this.start),this.$emit("goToStep",this.step)},updateQuery(e,d){const t=this.$router,u=this.$route;u.query?t.push({query:Object.assign({},u.query,{[e]:d})}):t.push({query:{[e]:d}})},setStartingStep(){const e=this.$route.query.step;this.start=e||0}}},x={class:"wizard-steps"},F={class:"wizard-steps__content-wrapper"},R={class:"wizard-steps__indicator"},N={class:"wizard-steps__indicator__controls",role:"tablist","aria-label":"steptabs"},T=["aria-selected","aria-controls"],A={class:"wizard-steps__content"},K={ref:"wizardForm",autocomplete:"off"},D=["id","aria-labelledby"],V={key:0,class:"wizard-steps__footer"},L={class:"wizard-steps__sidebar"},P={class:"wizard-steps__sidebar__content"};function Q(e,d,t,u,a,_){const p=f("KButton");return s(),n("div",x,[o("div",F,[o("header",R,[o("ul",N,[(s(!0),n(v,null,S(t.steps,(i,l)=>(s(),n("li",{key:i.slug,"aria-selected":_.step===i.slug?"true":"false","aria-controls":`wizard-steps__content__item--${l}`,class:m([{"is-complete":l<=a.start},"wizard-steps__indicator__item"])},[o("span",null,w(i.label),1)],10,T))),128))])]),o("div",A,[o("form",K,[(s(!0),n(v,null,S(t.steps,(i,l)=>(s(),n("div",{id:`wizard-steps__content__item--${l}`,key:i.slug,"aria-labelledby":`wizard-steps__content__item--${l}`,role:"tabpanel",tabindex:"0",class:"wizard-steps__content__item"},[_.step===i.slug?c(e.$slots,i.slug,{key:0},void 0,!0):r("",!0)],8,D))),128))],512)]),t.footerEnabled?(s(),n("footer",V,[b(h(p,{appearance:"outline",onClick:_.goToPrevStep},{default:y(()=>[C(" \u2039 Previous ")]),_:1},8,["onClick"]),[[k,!_.indexCanReverse]]),b(h(p,{disabled:t.nextDisabled,appearance:"primary",onClick:_.goToNextStep},{default:y(()=>[C(" Next \u203A ")]),_:1},8,["disabled","onClick"]),[[k,!_.indexCanAdvance]])])):r("",!0)]),o("aside",L,[o("div",P,[(s(!0),n(v,null,S(t.sidebarContent,(i,l)=>(s(),n("div",{key:i.name,class:m(["wizard-steps__sidebar__item",`wizard-steps__sidebar__item--${l}`])},[c(e.$slots,i.name,{},void 0,!0)],2))),128))])])])}const ne=g($,[["render",Q],["__scopeId","data-v-e6184caf"]]);const j={},U={class:"card-icon icon-success mb-3",role:"img"};function H(e,d){return s(),n("i",U," \u2713 ")}const O=g(j,[["render",H],["__scopeId","data-v-ec527321"]]);const G={name:"EntityScanner",components:{IconSuccess:O},props:{interval:{type:Number,required:!1,default:1e3},retries:{type:Number,required:!1,default:3600},shouldStart:{type:Boolean,default:!1},hasError:{type:Boolean,default:!1},loaderFunction:{type:Function,required:!0},canComplete:{type:Boolean,default:!1}},emits:["hide-siblings"],data(){return{i:0,isRunning:!1,isComplete:!1,intervalId:null}},watch:{shouldStart(e,d){e!==d&&e===!0&&this.runScanner()}},mounted(){this.shouldStart===!0&&this.runScanner()},beforeUnmount(){clearInterval(this.intervalId)},methods:{runScanner(){this.isRunning=!0,this.isComplete=!1,this.intervalId=setInterval(()=>{this.i++,this.loaderFunction(),(this.i===this.retries||this.canComplete===!0)&&(clearInterval(this.intervalId),this.isRunning=!1,this.isComplete=!0,this.$emit("hide-siblings",!0))},this.interval)}}},J={key:0,class:"scanner"},M={class:"scanner-content"},W={key:0,class:"card-icon mb-3"},X={key:1,class:"card-icon mb-3"},Y={key:3},Z={key:1};function ee(e,d,t,u,a,_){const p=f("KIcon"),i=f("IconSuccess"),l=f("KEmptyState");return t.shouldStart?(s(),n("div",J,[o("div",M,[h(l,{"cta-is-hidden":""},{title:y(()=>[a.isRunning?(s(),n("div",W,[h(p,{icon:"spinner",color:"rgba(0, 0, 0, 0.1)",size:"42"})])):r("",!0),a.isComplete&&t.hasError===!1&&a.isRunning===!1?(s(),n("div",X,[h(i)])):r("",!0),a.isRunning?c(e.$slots,"loading-title",{key:2},void 0,!0):r("",!0),a.isRunning===!1?(s(),n("div",Y,[t.hasError?c(e.$slots,"error-title",{key:0},void 0,!0):r("",!0),a.isComplete&&t.hasError===!1?c(e.$slots,"complete-title",{key:1},void 0,!0):r("",!0)])):r("",!0)]),message:y(()=>[a.isRunning?c(e.$slots,"loading-content",{key:0},void 0,!0):r("",!0),a.isRunning===!1?(s(),n("div",Z,[t.hasError?c(e.$slots,"error-content",{key:0},void 0,!0):r("",!0),a.isComplete&&t.hasError===!1?c(e.$slots,"complete-content",{key:1},void 0,!0):r("",!0)])):r("",!0)]),_:3})])])):r("",!0)}const re=g(G,[["render",ee],["__scopeId","data-v-5ffa385d"]]);export{re as E,se as F,ne as S};