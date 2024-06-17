import{d as L,r as a,w as v,h as k,c as s,u as l,a as e,e as _,i as V,j as b,f as p,t as x,g as f,k as g,F as C,l as h,o as n}from"./index-ae5a3ea2.js";const E={key:0},N=e("h1",null,"Loading...",-1),S=[N],j={key:1},R=e("h1",null,"Login for the name game!",-1),T=e("p",null,"Enter your email and password to login, if you do not have a username and password yet, register first",-1),B={key:0},D={class:"important"},P={key:0,class:"important-content green"},U=e("br",null,null,-1),q={key:1,class:"important-content red"},F=e("p",null,"Login failed!",-1),O=e("label",{for:"email"},"Email:",-1),A=e("label",{for:"password"},"Password:",-1),I="/api/names",$=L({__name:"LoginView",setup(J){const m=a(null);v(async()=>{const i=`${I}`;m.value=await(await fetch(i)).json()});const u=a(""),r=a("");var t=a(""),c=a("");v(()=>{t.value=""});function y(){const i={email:u.value,password:r.value};fetch("/api/voters/login",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(i)}).then(o=>o.json()).then(o=>{console.log("Success:",o),o.email?(console.log("Login successful"),t.value="success",h.push({path:"/voting"}).then(()=>{h.go(0)})):(console.log("Login failed"),t.value="failure",c.value=o.message)}).catch(o=>{console.error("Error:",o),t.value="failure",c.value=o,alert("Error submitting vote: "+o)})}return(i,o)=>{const w=k("RouterLink");return n(),s(C,null,[m.value?(n(),s("div",j,[R,T,l(t)==="success"||l(t)==="failure"?(n(),s("div",B,[e("div",D,[l(t)==="success"?(n(),s("div",P,[e("p",null,[_("Login successful! "),U,_("Now you can "),V(w,{to:"/voting"},{default:b(()=>[_("Vote!")]),_:1})])])):p("",!0),l(t)==="failure"?(n(),s("div",q,[F,e("p",null,"Details: "+x(l(c)),1)])):p("",!0)])])):p("",!0)])):(n(),s("div",E,S)),e("div",null,[e("form",null,[O,f(e("input",{type:"email",id:"email","onUpdate:modelValue":o[0]||(o[0]=d=>u.value=d),required:""},null,512),[[g,u.value]]),A,f(e("input",{type:"password",id:"password","onUpdate:modelValue":o[1]||(o[1]=d=>r.value=d),required:""},null,512),[[g,r.value]]),e("button",{type:"submit",onClick:y},"Login")])])],64)}}});export{$ as default};
