import{d as T,r as c,w as _,c as t,a as e,t as l,F as f,b as g,e as b,u as d,f as v,o as s,g as S,v as x}from"./index-74911e70.js";const j={key:0},N=e("h1",null,"Loading...",-1),C=[N],R={key:1},E=e("h1",null,"Register and Login to bet 🥳",-1),L=e("p",null,"After you are logged-in you can place your bets!",-1),P=e("br",null,null,-1),Y={style:{"font-style":"italic"}},$={key:2},A={key:0},B=e("h1",null,"Thank you for voting 💙",-1),D=e("br",null,null,-1),F=e("p",null,"Transfer the bet money to one of the following accounts to finalize your participation:",-1),I=e("td",{class:"label"},"Amount",-1),O={class:"value"},U=e("td",{class:"label"},"Twint",-1),z={class:"value"},J=e("td",{class:"label"},"Revolut",-1),M={class:"value"},W=e("td",{class:"label"},"Paypal",-1),q={class:"value"},G=e("br",null,null,-1),H=e("p",null,"Thank you for participating in the name game!",-1),K=e("p",null,"We will be sending out the results after the 20.08.2024 🥰",-1),Q=e("br",null,null,-1),X=e("h2",null,"Your votes:",-1),Z={key:1},ee=e("h1",null,"Place your guess 🍼",-1),te=e("p",null,"Select all your guesses, but choose wisely. You can only guess once!",-1),se=e("p",null,"The vote only really counts once you transferred the matching sum",-1),le=e("hr",null,null,-1),oe=e("br",null,null,-1),ne=["value"],ae={class:"message"},ue=e("br",null,null,-1),ie=e("br",null,null,-1),ce=e("hr",null,null,-1),re=e("br",null,null,-1),_e={key:0},de={class:"important"},he={key:0,class:"important-content green"},ve=e("p",null,"Voting successful!",-1),pe=[ve],ye={key:1,class:"important-content red"},me=e("p",null,"Voting has failed!",-1),fe={key:1},ge=e("p",null,"You will be voting for:",-1),be={style:{"font-weight":"bold","text-decoration":"underline"}},we=e("p",null,"You can use twint, revolut or paypal - please mark your name when transferring 💙",-1),ke=e("br",null,null,-1),p="/api",Se=T({__name:"VotingView",setup(Ve){const r=c(null);var u=c(""),y=c("");const w=c(""),h=c(""),m=c([]);_(async()=>{const o=`${p}/me`;w.value=await(await fetch(o)).json()}),_(async()=>{const o=`${p}/votes/me`;m.value=await(await fetch(o)).json()}),_(async()=>{const o=`${p}/payments`;h.value=await(await fetch(o)).json()}),_(()=>{u.value="",y.value=""}),_(async()=>{const o=`${p}/names`;r.value=await(await fetch(o)).json()});const i=c([]);function k(){console.log("Submitting vote for",i.value);var o=i.value;fetch("/api/votes",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({names:o})}).then(n=>n.json()).then(n=>{console.log("Success:",n),n.email?(u.value="success",setTimeout(()=>{location.reload()},2e3)):(u.value="failure",y.value=n.message)}).catch(n=>{console.error("Error:",n),alert("Error submitting vote")})}return(o,n)=>r.value?r.value.message?(s(),t("div",R,[E,L,P,e("p",Y,"Details: "+l(r.value.message),1)])):(s(),t("div",$,[w.value.has_voted?(s(),t("div",A,[B,D,F,e("table",null,[e("tr",null,[I,e("td",O,l(m.value.length*10)+".00",1)]),e("tr",null,[U,e("td",z,l(h.value.twint),1)]),e("tr",null,[J,e("td",M,l(h.value.revolut),1)]),e("tr",null,[W,e("td",q,l(h.value.paypal),1)])]),G,H,K,Q,X,e("ul",null,[(s(!0),t(f,null,g(m.value,a=>(s(),t("li",{key:a},l(a),1))),128))])])):(s(),t("div",Z,[ee,te,e("p",null,[b("There are "),e("b",null,l(r.value.names.length),1),b(" available names to choose from.")]),se,le,oe,e("ul",null,[(s(!0),t(f,null,g(r.value.names,a=>(s(),t("p",{key:a.value},[S(e("input",{type:"checkbox","onUpdate:modelValue":n[0]||(n[0]=V=>i.value=V),value:a},null,8,ne),[[x,i.value]]),e("span",ae,"   "+l(a),1),ue]))),128))]),ie,ce,re,d(u)==="success"||d(u)==="failure"?(s(),t("div",_e,[e("div",de,[d(u)==="success"?(s(),t("div",he,pe)):v("",!0),d(u)==="failure"?(s(),t("div",ye,[me,e("p",null,"Reason: "+l(d(y)),1)])):v("",!0)])])):v("",!0),i.value.length>0?(s(),t("div",fe,[ge,e("ul",null,[(s(!0),t(f,null,g(i.value,a=>(s(),t("li",{key:a},l(a),1))),128))]),e("p",null,[b("This results in a bet of "),e("span",be,l(i.value.length*10)+" Franken",1)]),we])):v("",!0),ke,e("button",{onClick:k,class:"submit-button"},"Submit Vote")]))])):(s(),t("div",j,C))}});export{Se as default};
