import{d as T,r as c,w as r,c as t,a as e,t as o,F as m,b as g,e as b,u as d,f as p,o as s,g as S,v as x}from"./index-ad4fb6a5.js";const j={key:0},N=e("h1",null,"Loading...",-1),C=[N],R={key:1},Y=e("h1",null,"Register and Login to bet 🥳",-1),E=e("p",null,"After you are logged-in you can place your bets!",-1),L=e("br",null,null,-1),P={style:{"font-style":"italic"}},$={key:2},A={key:0},B=e("h1",null,"Thank you for voting 💙",-1),D=e("br",null,null,-1),F=e("p",null,"Transfer the bet money to one of the following accounts to finalize your participation:",-1),I=e("td",{class:"label"},"Amount",-1),O={class:"value"},U=e("td",{class:"label"},"Twint",-1),z={class:"value"},J=e("td",{class:"label"},"Revolut",-1),M={class:"value"},W=e("td",{class:"label"},"Paypal",-1),q={class:"value"},G=e("br",null,null,-1),H=e("p",null,"Thank you for participating in the name game!",-1),K=e("p",null,"We will be sending out the results after the 20.08.2024 🥰",-1),Q=e("br",null,null,-1),X=e("h2",null,"Your votes:",-1),Z={key:1},ee=e("h1",null,"Place your guess 🍼",-1),te=e("p",null,"Select all your guesses, but choose wisely. You can only guess once!",-1),se=e("p",null,"The vote only really counts once you transferred the matching sum",-1),le=e("p",null,"You will get the payment information after you placed your votes",-1),oe=e("hr",null,null,-1),ne=["value"],ae={class:"message"},ue=e("br",null,null,-1),ie=e("br",null,null,-1),ce=e("hr",null,null,-1),_e=e("br",null,null,-1),re={key:0},de={class:"important"},he={key:0,class:"important-content green"},ve=e("p",null,"Voting successful!",-1),pe=[ve],ye={key:1,class:"important-content red"},fe=e("p",null,"Voting has failed!",-1),me={key:1},ge=e("p",null,"You will be voting for:",-1),be={style:{"font-weight":"bold","text-decoration":"underline"}},we=e("p",null,"You can use twint, revolut or paypal - please mark your name when transferring 💙",-1),ke=e("br",null,null,-1),y="/api",Se=T({__name:"VotingView",setup(Ve){const _=c(null);var u=c(""),h=c("");const w=c(""),v=c(""),f=c([]);r(async()=>{const n=`${y}/me`;w.value=await(await fetch(n)).json()}),r(async()=>{const n=`${y}/votes/me`;f.value=await(await fetch(n)).json()}),r(async()=>{const n=`${y}/payments`;v.value=await(await fetch(n)).json()}),r(()=>{u.value="",h.value=""}),r(async()=>{const n=`${y}/names`;_.value=await(await fetch(n)).json()});const i=c([]);function k(){console.log("Submitting vote for",i.value);var n=i.value;fetch("/api/votes",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({names:n})}).then(l=>l.json()).then(l=>{console.log("Success:",l),l.email?(u.value="success",setTimeout(()=>{location.reload()},2e3)):(u.value="failure",h.value=l.message)}).catch(l=>{console.error("Error:",l),u.value="failure",h.value=l,alert("Error submitting vote: "+l)})}return(n,l)=>_.value?_.value.message?(s(),t("div",R,[Y,E,L,e("p",P,"Details: "+o(_.value.message),1)])):(s(),t("div",$,[w.value.has_voted?(s(),t("div",A,[B,D,F,e("table",null,[e("tr",null,[I,e("td",O,o(f.value.length*10)+".00",1)]),e("tr",null,[U,e("td",z,o(v.value.twint),1)]),e("tr",null,[J,e("td",M,o(v.value.revolut),1)]),e("tr",null,[W,e("td",q,o(v.value.paypal),1)])]),G,H,K,Q,X,e("ul",null,[(s(!0),t(m,null,g(f.value,a=>(s(),t("li",{key:a},o(a),1))),128))])])):(s(),t("div",Z,[ee,te,e("p",null,[b("There are "),e("b",null,o(_.value.names.length),1),b(" available names to choose from.")]),se,le,oe,e("ul",null,[(s(!0),t(m,null,g(_.value.names,a=>(s(),t("p",{key:a.value},[S(e("input",{type:"checkbox","onUpdate:modelValue":l[0]||(l[0]=V=>i.value=V),value:a},null,8,ne),[[x,i.value]]),e("span",ae,"   "+o(a),1),ue]))),128))]),ie,ce,_e,d(u)==="success"||d(u)==="failure"?(s(),t("div",re,[e("div",de,[d(u)==="success"?(s(),t("div",he,pe)):p("",!0),d(u)==="failure"?(s(),t("div",ye,[fe,e("p",null,"Reason: "+o(d(h)),1)])):p("",!0)])])):p("",!0),i.value.length>0?(s(),t("div",me,[ge,e("ul",null,[(s(!0),t(m,null,g(i.value,a=>(s(),t("li",{key:a},o(a),1))),128))]),e("p",null,[b("This results in a bet of "),e("span",be,o(i.value.length*10)+" Franken",1)]),we])):p("",!0),ke,e("button",{onClick:k,class:"submit-button"},"Submit Vote")]))])):(s(),t("div",j,C))}});export{Se as default};
