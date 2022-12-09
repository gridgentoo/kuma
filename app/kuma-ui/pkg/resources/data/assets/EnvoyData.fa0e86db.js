import{d as y,O as p,o as r,j as m,c as l,J as g,r as i,g as d,h,k as u,w as _,a as v,E}from"./index.5ae08964.js";import{_ as N}from"./CodeBlock.vue_vue_type_style_index_0_lang.c46d061b.js";import{_ as k}from"./EmptyBlock.vue_vue_type_script_setup_true_lang.9cbf2f16.js";import{E as q}from"./ErrorBlock.e3550eb2.js";import{_ as z}from"./LoadingBlock.vue_vue_type_script_setup_true_lang.91929754.js";const P={key:3},S=y({__name:"StatusInfo",props:{isLoading:{type:Boolean,default:!1},hasError:{type:Boolean,default:!1},isEmpty:{type:Boolean,default:!1},error:{type:[Error,p],required:!1,default:null}},setup(t){return(e,s)=>(r(),m("div",null,[t.isLoading?(r(),l(z,{key:0})):t.hasError?(r(),l(q,{key:1,error:t.error},null,8,["error"])):t.isEmpty?(r(),l(k,{key:2})):(r(),m("div",P,[g(e.$slots,"default")]))]))}}),I=y({__name:"EnvoyData",props:{dataPath:{type:String,required:!0},queryKey:{type:String,required:!1,default:null},mesh:{type:String,required:!1,default:""},dppName:{type:String,required:!1,default:""},zoneIngressName:{type:String,required:!1,default:""},zoneEgressName:{type:String,required:!1,default:""}},setup(t){const e=t,s=i(!0),o=i(void 0),c=i("");d(()=>e.dppName,function(){n()}),d(()=>e.zoneIngressName,function(){n()}),d(()=>e.zoneEgressName,function(){n()}),h(function(){n()});async function n(){o.value=void 0,s.value=!0;try{let a="";e.mesh!==""&&e.dppName!==""?a=await u.getDataplaneData({dataPath:e.dataPath,mesh:e.mesh,dppName:e.dppName}):e.zoneIngressName!==""?a=await u.getZoneIngressData({dataPath:e.dataPath,zoneIngressName:e.zoneIngressName}):e.zoneEgressName!==""&&(a=await u.getZoneEgressData({dataPath:e.dataPath,zoneEgressName:e.zoneEgressName})),c.value=typeof a=="string"?a:JSON.stringify(a,null,2)}catch(a){a instanceof Error?o.value=a:console.error(a)}finally{s.value=!1}}return(a,B)=>(r(),l(S,{class:"envoy-data","has-error":o.value!==void 0,"is-loading":s.value,error:o.value},{default:_(()=>{var f;return[v(N,{id:`code-block-${t.dataPath}`,language:"json",code:c.value,"is-searchable":"","query-key":(f=t.queryKey)!=null?f:`code-block-${t.dataPath}`},null,8,["id","code","query-key"])]}),_:1},8,["has-error","is-loading","error"]))}});const C=E(I,[["__scopeId","data-v-b9869c64"]]);export{C as E,S as _};
