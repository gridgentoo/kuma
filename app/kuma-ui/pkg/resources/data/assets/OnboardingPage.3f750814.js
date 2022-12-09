import{E as r,o,j as d,l as t,t as c,e as a,A as g,J as s,B as u,C as h,D as b}from"./index.5ae08964.js";const m={name:"OnboardingHeading",props:{title:{type:String,required:!0},description:{type:String,default:""}}},v={class:"relative"},f={class:"onboarding-title"},$={key:0,class:"text-center text-lg mt-3"};function y(e,_,n,l,p,i){return o(),d("div",v,[t("h1",f,c(n.title),1),a(),n.description?(o(),d("p",$,c(n.description),1)):g("",!0)])}const H=r(m,[["render",y],["__scopeId","data-v-8a9a6bad"]]);const I={name:"OnboardingContainer",props:{withImage:{type:Boolean,default:!1}},computed:{classes(){return["onboarding-container__content",this.withImage?"onboarding-container__content--with-image":""]}}},O=e=>(h("data-v-0bedcde9"),e=e(),b(),e),S={class:"onboarding-container"},x={class:"onboarding-container__header"},w={class:"w-full"},B=O(()=>t("div",{class:"background-image"},null,-1));function k(e,_,n,l,p,i){return o(),d("div",null,[t("div",S,[t("div",x,[s(e.$slots,"header",{},void 0,!0)]),a(),t("div",{class:u(i.classes)},[t("div",w,[s(e.$slots,"content",{},void 0,!0)])],2),a(),s(e.$slots,"navigation",{},void 0,!0)]),a(),B])}const N=r(I,[["render",k],["__scopeId","data-v-0bedcde9"]]);export{H as O,N as a};
