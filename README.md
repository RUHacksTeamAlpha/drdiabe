# drdiabe.tech

# 🔗Important Links 🔗
Deployed instance: https://drdiabe.tech 

API: https://api.drdiabe.tech/ping

Slides: https://docs.google.com/presentation/d/19DkmZtyuD7ThfD9csdANWjIfbcv20jqQXNovqAvv7tA/edit?usp=sharing

# ✨ Inspiration ✨
We were inspired the millions of people with diabetes around the world that are unable to afford or access top-of-the-line equipment. Specifically, we identified the pain that the archaic technology presented and looked to disrupt current practices with a scalable web app.

# ⚒ What it does ⚒
DrDiabeTech provides users the ability to enter diabetic information to be stored in the cloud. The platform then enables rich insight into how users are managing their condition through data visualization and trend analysis. The system was built out of the box to be scalable and is fully deployed for use today!

# 👷‍♂️ How we built it 👷‍♀️
The platform was built using Vue.js on the Nuxt.js framework on the front-end. This system makes it easy to rapidly develop an app with data visualization tools, without compromise on speed or SEO ability (thanks the the server side rendering).

The backend for logic was built in Golang with help from the Gin framework. Our Golang JSON REST API communicates to the frontend to send and receive all user information.

User accounts are created and managed by Google Firebase, allowing for out-of-the-box verification and account resetting. As well, firebase serves to authenticate user requests by providing the frontend with a token upon log-in that can be passed to the Golang backend and the once again verified against the firebase servers.

User data is saved in Google Cloud SQL to allow a fast, yet secure and reliable persistent data store. This is useful to people who may access the web app from a variety of devices. As well, the information can be downloaded by the user to perform their own analysis.

The last part of the application is deployment which is handled by using Docker and Google Cloud Run. A single GCR service supports up to 100 000 users at maximum scale, thus our app can already be used by thousands of people across the world.

# 💪 Challenges we ran into 💪
As we were only two people, it was tough to work on such a wide web app. We had to refocus our efforts a number of times to ensure we could fully complete the project instead of leaving many features not completed.

# 🥇 Accomplishments that we're proud of 🥇
We are super proud that the project is deployed and accessible for anyone to use! As well, we were able to apply our knowledge of SWE that we learned in university through using tools to test the accessibility of the web site and properly access use cases.

# 🚸 What we learned 🚸
We learned a lot of nitty gritty skills for managing data. As well, working with decoupled technologies was a good lesson as to how people working in different areas must communicate. 

# 😲 What's next for DrDiabe.Tech😲
We have more ideas on analyzing negative trends and alerting users. As well, we wanted to implement the ability to take a picture of an insulin meter and use OCR to automatically record the info, but ran out of time
