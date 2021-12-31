var L=Object.defineProperty,A=Object.defineProperties;var O=Object.getOwnPropertyDescriptors;var m=Object.getOwnPropertySymbols;var T=Object.prototype.hasOwnProperty,j=Object.prototype.propertyIsEnumerable;var y=(r,e,t)=>e in r?L(r,e,{enumerable:!0,configurable:!0,writable:!0,value:t}):r[e]=t,d=(r,e)=>{for(var t in e||(e={}))T.call(e,t)&&y(r,t,e[t]);if(m)for(var t of m(e))j.call(e,t)&&y(r,t,e[t]);return r},a=(r,e)=>A(r,O(e));import{s,r as x,R as v,j as p,a as o,u as b,b as C,c as I,d as P,e as R,f as k,t as q,g as F,P as M}from"./vendor.55af4f09.js";const N=function(){const e=document.createElement("link").relList;if(e&&e.supports&&e.supports("modulepreload"))return;for(const i of document.querySelectorAll('link[rel="modulepreload"]'))n(i);new MutationObserver(i=>{for(const l of i)if(l.type==="childList")for(const c of l.addedNodes)c.tagName==="LINK"&&c.rel==="modulepreload"&&n(c)}).observe(document,{childList:!0,subtree:!0});function t(i){const l={};return i.integrity&&(l.integrity=i.integrity),i.referrerpolicy&&(l.referrerPolicy=i.referrerpolicy),i.crossorigin==="use-credentials"?l.credentials="include":i.crossorigin==="anonymous"?l.credentials="omit":l.credentials="same-origin",l}function n(i){if(i.ep)return;i.ep=!0;const l=t(i);fetch(i.href,l)}};N();const G=s.div`
  background-color: #ededed;
  width: 100vw;
  height: 100vh;
  overflow:hidden;
  display:grid;
  place-items:center;
`,U=s.div`
  width: 40vw;
  height: 80vh;
  overflow: hidden;
  display:flex;
  flex-direction: column;
  background-color: white;
  box-shadow: 0px 3px 30px rgba(0,0,0,0.2);
  border-radius: 10px;
  margin-right: 5px;
`,H=s.div`
  background-color: lightgreen;
  color: #323232;
  font-weight: bold;
  padding: 1.2em;
`,z=s.div`
  display:flex;
  flex-direction: row;
`,B=s.form`
  display: flex;
  flex-direction: column;
  margin: 5px;
  margin-left: 15px;
  background-color: white;
  box-shadow: 0px 3px 30px rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  padding: 0px 1em;
`,u=s.input`
  margin-bottom: 5px;
  border: 1px solid gray;
  border-radius: 5px;
  padding: 1em;
  font-size: 1em;
  outline: none;
`;let w={name:"",address:"",image:""};const J=({addSupplier:r})=>{const[e,t]=x.exports.useState(w),n=v.useRef(null);return p(B,{onSubmit:i=>{if(i.preventDefault(),e.name.length===0&&e.address.length===0&&e.image.length===0){alert("Please fill name, address and select a image to add a supplier");return}r(e),t(w),n.current&&(n.current.value="")},children:[o("h3",{children:"Add new supplier"}),o(u,{required:!0,placeholder:"Type supplier name here",value:e.name,type:"text",onChange:i=>t(l=>a(d({},l),{name:i.target.value}))}),o(u,{ref:n,required:!0,placeholder:"select supplier photo here",accept:"image/jpeg, image/png",type:"file",onChange:i=>{let l=i.target.files?i.target.files[0]:"";t(c=>a(d({},c),{image:l}))}}),o(u,{required:!0,placeholder:"Type supplier address here",value:e.address,type:"text",onChange:i=>t(l=>a(d({},l),{address:i.target.value}))}),o(u,{type:"submit",value:"Submit"})]})},K=r=>fetch("/api/uploadImage",{method:"POST",body:r}).then(e=>e.json()).then(e=>e.data),h=({url:r,method:e,data:t})=>{let n={method:e,headers:{"Content-Type":"application/json"}};return e!=="GET"?fetch(r,a(d({},n),{body:JSON.stringify(t)})).then(i=>i.json()):fetch(r,n).then(i=>i.json())},f={get:()=>h({url:"/api/suppliers",method:"GET",data:{}}).then(({data:r})=>r),add:async(r,e)=>{const{data:t}=await h({url:"/api/suppliers",method:"POST",data:r});e.append("sup_id",t.id);const n=await K(e);return a(d({},t),{image:n})},edit:r=>h({url:"/api/suppliers",method:"PUT",data:r}).then(({data:e})=>e),delete:r=>h({url:"/api/suppliers",method:"DELETE",data:r}).then(({data:e})=>e)},S="supplier/add",D="supplier/del",E="supplier/get",g="supplier/loading",W=()=>r=>{r({type:g}),f.get().then(e=>r({type:E,payload:e}))},_=r=>(e,t)=>{t().suppliers,e({type:g}),f.delete(r).then(n=>{e({type:D,payload:n==="Done"?r.id:-1})})},$=(r,e)=>t=>{t({type:g}),f.add(r,e).then(n=>t({type:S,payload:n}))},Q=s.div`
    display:flex;
    flex-direction: row;
    align-items: center;
    margin-top: 10px;
    border-bottom: 1px solid lightgray;
    padding-bottom: 5px;
`,V=s.img`
    width: 48px;
    height: 48px;
    border-radius:; 48px;
    padding-right: 10px;
`,X=({supplier:r,onDelete:e})=>p(Q,{onClick:e,children:[o(V,{src:r.image}),p("div",{style:{display:"flex",flex:1,flexDirection:"column"},children:[o("div",{children:r.name}),o("div",{children:r.address})]}),o("div",{style:{padding:5,cursor:"pointer",borderRadius:2,backgroundColor:"crimson",color:"white"},children:"delete"})]}),Y=({suppliers:r})=>{const e=b(),t=n=>{e(_(n))};return o("div",{style:{overflow:"scroll",padding:"0px 1em"},children:r.map(n=>o(X,{supplier:n,onDelete:()=>t(n)},n.id+"Hello"))})};function Z(){const{suppliers:r}=C(n=>n.suppliers,I),e=b();x.exports.useEffect(()=>{e(W())},[]);const t=n=>{let i=new FormData;i.append("image",n.image),e($({name:n.name,address:n.address},i))};return o(G,{children:p(z,{children:[p(U,{children:[o(H,{children:p("strong",{children:["Suppliers [",r.length,"]"]})}),o(Y,{suppliers:r})]}),o(J,{addSupplier:n=>t(n)})]})})}const ee={suppliers:[],loading:!1},re=(r=ee,e)=>{switch(e.type){case E:return a(d({},r),{suppliers:e.payload,loading:!1});case S:return a(d({},r),{loading:!1,suppliers:[...r.suppliers,e.payload]});case D:return a(d({},r),{loading:!1,suppliers:r.suppliers.filter(t=>t.id!==e.payload)});case g:return a(d({},r),{loading:!0});default:return r}},te=P({suppliers:re});var ie=R(te,k(q));F.render(o(v.StrictMode,{children:o(M,{store:ie,children:o(Z,{})})}),document.getElementById("root"));
