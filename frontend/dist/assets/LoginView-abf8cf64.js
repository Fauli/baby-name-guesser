import{d as y,r as i,w as m,h as k,c as t,u as l,a as e,b as d,i as L,j as V,f as _,g as v,k as h,F as b,l as f,o as n}from"./index-f956925b.js";const x={key:0},C=e("h1",null,"Loading...",-1),N=[C],E={key:1},T=e("h1",null,"Login for the name game!",-1),j={key:0},B={class:"important"},P={key:0,class:"important-content green"},R=e("br",null,null,-1),S={key:1,class:"important-content red"},U=e("p",null,"Login failed!",-1),q=[U],F=e("label",{for:"email"},"Email:",-1),O=e("label",{for:"password"},"Password:",-1),A="/api/names",J=y({__name:"LoginView",setup(D){const p=i(null);m(async()=>{const a=`${A}`;p.value=await(await fetch(a)).json()});const r=i(""),c=i("");var o=i("");m(()=>{o.value=""});const g=async()=>{const a={email:r.value,password:c.value};try{(await fetch("/api/voters/login",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(a)})).ok?(console.log("Login successful"),o.value="success",f.push({path:"/voting"}).then(()=>{f.go(0)})):(console.error("Login failed"),o.value="failure")}catch(s){console.error("Error:",s)}};return(a,s)=>{const w=k("RouterLink");return n(),t(b,null,[p.value?(n(),t("div",E,[T,l(o)==="success"||l(o)==="failure"?(n(),t("div",j,[e("div",B,[l(o)==="success"?(n(),t("div",P,[e("p",null,[d("Login successful! "),R,d("Now you can "),L(w,{to:"/voting"},{default:V(()=>[d("Vote!")]),_:1})])])):_("",!0),l(o)==="failure"?(n(),t("div",S,q)):_("",!0)])])):_("",!0)])):(n(),t("div",x,N)),e("div",null,[e("form",null,[F,v(e("input",{type:"email",id:"email","onUpdate:modelValue":s[0]||(s[0]=u=>r.value=u),required:""},null,512),[[h,r.value]]),O,v(e("input",{type:"password",id:"password","onUpdate:modelValue":s[1]||(s[1]=u=>c.value=u),required:""},null,512),[[h,c.value]]),e("button",{onClick:g},"Login")])])],64)}}});export{J as default};
