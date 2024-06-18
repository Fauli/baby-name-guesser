<script setup lang="ts">
import { ref } from 'vue';

const isVisible = ref(false);

const toggleVisibility = () => {
  isVisible.value = !isVisible.value;
}

</script>

<template>
    <header>
        <h1>BaNaGue - Baby Name Guesser!</h1>
    </header>
    <br />
    <section>
        <p>Welcome to our Baby Name Guessing Game!!! 👶</p>
        <p>Für eine deutsche Version klicke den Knopf ganz unten</p>
    </section>
    <br />
    <section>
        <p>Exciting times await as we prepare to welcome our newest family member! 🐣<br />
            But before the big reveal, we invite you to join in a fun and unique challenge: <span
                style="font-weight: bold;">Guess our baby’s name!</span><br />
            Participate, have fun, and contribute to our little one's future all at once. 💙</p>
    </section>
    <br />
    <section>
        <h2>How It Works</h2>
        <ul>
            <li>We’ve shortlisted some beautiful names for our baby. Your challenge is to guess the one we’ve chosen as the
                <span style="font-weight: bold;">first name.</span>
            </li>
            <li>Place your votes — each guess costs CHF 10, and you can cast as many guesses as you like to increase your
                chances of winning, but only once... so choose wisely!</li>
            <li>Every single Franken raised will go into a fund for our baby’s future, making each guess a gift to our
                child.</li>
            <li>Guess correctly, and you’ll win a treat — a delicious three-course meal 🍲 prepared by us! The date of the meal will be arranged after her birth.</li>
        </ul>
    </section>
    <br />
    <section>
        <h2>Important Rules</h2>
        <ul>
            <li>Make your choice carefully. Once your vote is cast and the payment is made through Twint, PayPal, IBAN, or
                Revolut, there's no changing it.</li>
            <li>After voting, you will be able to see the overall top 3 guessed names.</li>
            <li>We already made the final decision on the name - so the game will not influence our choice 😁</li>
        </ul>
    </section>
    <br />
    <section>
        <p>Join us in this joyful prelude to our baby's arrival and possibly earn yourself a delightful dinner. Ready to
            play? Cast your vote and let the guessing begin!</p>
    </section>

    <br />
    <br />
    <button @click="toggleVisibility" class="toggle-button">Deutsche Version anzeigen</button>
    <div v-show="isVisible">
        <br />
        <br />

        <section>
            <p>Willkommen bei unserem Baby Namen Ratespiel!!! 👶</p>
        </section>
        <br />
        <section>
            <p>Aufregende Zeiten kommen auf uns zu und wir bereiten uns darauf vor, unser neuestes Familienmitglied
                willkommen zu heissen! 🐣<br />
                Aber bevor wir das grosse Namensgeheimnis lüften, laden wir dich ein, an einer lustigen Herausforderung
                teilzunehmen: <span style="font-weight: bold;">Rate den Namen unseres
                    Babys!</span><br />
                Mach mit, hab Spass und trage gleichzeitig zur Zukunft unserer Kleinen bei. 💙</p>
        </section>
        <br />
        <section>
            <h2>Wie es funktioniert</h2>
            <ul>
                <li>Wir haben einige schöne Namen für unser Baby zusammengetragen. Deine Herausforderung besteht darin,
                    den Namen zu erraten, den wir als <span style="font-weight: bold;">ersten Vornamen</span> gewählt
                    haben.</li>
                <li>Gib deine Stimmen ab — jeder Tipp kostet CHF 10, und du kannst so viele Tipps abgeben, wie du
                    möchtest, um deine Gewinnchancen zu erhöhen, aber nur einmal... also wähle weise!</li>
                <li>Jeder einzelne Franken, der gesammelt wird, fliesst in einen Fonds für die Zukunft unseres Babys,
                    sodass jeder Tipp ein Geschenk für unser Kind ist.</li>
                <li>Errätst du den Namen richtig, erwartet dich eine Belohnung — ein exklusives Drei-Gang-Menü 🍲, das von uns
                    zubereitet wird! Der Termin für das Essen wird nach der Geburt vereinbart.</li>
            </ul>
        </section>
        <br />
        <section>
            <h2>Wichtige Regeln</h2>
            <ul>
                <li>Treffe deine Wahl sorgfältig. Sobald deine Stimme abgegeben ist, gibt es kein Zurück mehr.</li>
                <li>Bezahle deinen Einsatz via Twint, PayPal, IBAN oder Revolut gemäss Angaben auf der Webseite, damit deine
                    Wahl auch vollständig gültig wird.</li>
                <li>Nachdem du abgestimmt hast, kannst du die drei am häufigsten gewählten Namen ansehen.</li>
                <li>Unsere Entscheidung ist bereits gefallen - das Spiel wird unsere Namenswahl also nicht beeinflussen 😁</li>
            </ul>
        </section>
        <br />
        <section>
            <p>Schliesse dich unserer freudigen Erwartung an und versuche dein Glück für ein exklusives Essen bei uns.
                Bereit zum Spielen? Registriere dich jetzt und lass das Raten beginnen!</p>
        </section>
    </div>

</template>


<style scoped>

.toggle-button {
  background-color: #36a7c6; /* Green background */
  border: none;
  color: white; /* White text */
  padding: 10px 20px; /* Some padding */
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  border-radius: 8px; /* Rounded corners */
  transition: background-color 0.3s, box-shadow 0.3s; /* Smooth transitions */
}

.toggle-button:hover {
  background-color: #1e85a2; /* Darker green background on hover */
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2); /* Add shadow on hover */
}

h2 {
    padding-bottom: 3px;
}

</style>