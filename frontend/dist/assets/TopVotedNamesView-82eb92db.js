import{d as p,r as a,w as i,c as o,a as e,t as _,F as f,e as v,o as n,m}from"./index-f956925b.js";const y={key:0},g=e("h1",null,"Loading...",-1),k=[g],w={key:1},x=e("h1",null,"Not quite yet...",-1),V=e("p",null,"You will only be able to view the top 3 guessed names so far after you guessed yourself ⭐️",-1),b=e("br",null,null,-1),N={style:{"font-style":"italic"}},B={key:2},L=e("h1",null,"Top 3 guesses 📊",-1),R=e("p",null,"Get a sneak peek on the 3 top guessed names so far...",-1),T=e("hr",null,null,-1),D={style:{"padding-left":"0px"}},E={style:{"font-size":"xx-large"},class:"message"},F=e("br",null,null,-1),S="/api/votes/top",A=p({__name:"TopVotedNamesView",setup(j){const t=a(null);var u=a(""),r=a(""),l=1;i(()=>{u.value="",r.value=""}),i(async()=>{const s=`${S}`;t.value=await(await fetch(s)).json()}),a([]);function c(s){return s==1?"🥇":s==2?"🥈":s==3?"🥉":"🏅"}return(s,q)=>t.value?t.value.message?(n(),o("div",w,[x,V,b,e("p",N,"Details: "+_(t.value.message),1)])):(n(),o("div",B,[L,R,T,e("ul",D,[(n(!0),o(f,null,v(t.value,(d,h)=>(n(),o("p",{key:h},[e("span",E,"   "+_(c(m(l)?l.value++:l++))+" "+_(d),1),F]))),128))])])):(n(),o("div",y,k))}});export{A as default};
