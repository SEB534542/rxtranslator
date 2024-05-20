package rxtranslator

import (
	//"os"
	"testing"
)

func TestResumeImportAndExport(t *testing.T) {
	data := []byte(`{
		"basics": {
		  "url": {
			"href": "https://johndoe.me/",
			"label": ""
		  },
		  "name": "John Doe",
		  "email": "john.doe@gmail.com",
		  "phone": "(555) 123-4567",
		  "picture": {
			"url": "https://i.imgur.com/HgwyOuJ.jpg",
			"size": 120,
			"effects": {
			  "border": false,
			  "hidden": false,
			  "grayscale": false
			},
			"aspectRatio": 1,
			"borderRadius": 0
		  },
		  "headline": "Creative and Innovative Web Developer",
		  "location": "Pleasantville, CA 94588",
		  "customFields": []
		},
		"metadata": {
		  "css": {
			"value": ".section {\n\toutline: 1px solid #000;\n\toutline-offset: 4px;\n}",
			"visible": false
		  },
		  "page": {
			"format": "a4",
			"margin": 14,
			"options": {
			  "breakLine": true,
			  "pageNumbers": true
			}
		  },
		  "notes": "",
		  "theme": {
			"text": "#000000",
			"primary": "#ca8a04",
			"background": "#ffffff"
		  },
		  "layout": [
			[
			  [
				"summary",
				"experience",
				"education",
				"projects",
				"references"
			  ],
			  [
				"profiles",
				"skills",
				"certifications",
				"interests",
				"languages",
				"awards",
				"volunteer",
				"publications"
			  ]
			]
		  ],
		  "template": "glalie",
		  "typography": {
			"font": {
			  "size": 13,
			  "family": "Merriweather",
			  "subset": "latin",
			  "variants": [
				"regular"
			  ]
			},
			"hideIcons": false,
			"lineHeight": 1.75,
			"underlineLinks": true
		  }
		},
		"sections": {
		  "awards": {
			"id": "awards",
			"name": "Awards",
			"items": [],
			"columns": 1,
			"visible": true
		  },
		  "custom": {},
		  "skills": {
			"id": "skills",
			"name": "Skills",
			"items": [
			  {
				"id": "hn0keriukh6c0ojktl9gsgjm",
				"name": "Web Technologies",
				"level": 0,
				"visible": true,
				"keywords": [
				  "HTML5",
				  "JavaScript",
				  "PHP",
				  "Python"
				],
				"description": "Advanced"
			  },
			  {
				"id": "r8c3y47vykausqrgmzwg5pur",
				"name": "Web Frameworks",
				"level": 0,
				"visible": true,
				"keywords": [
				  "React.js",
				  "Angular",
				  "Vue.js",
				  "Laravel",
				  "Django"
				],
				"description": "Intermediate"
			  },
			  {
				"id": "b5l75aseexqv17quvqgh73fe",
				"name": "Tools",
				"level": 0,
				"visible": true,
				"keywords": [
				  "Webpack",
				  "Git",
				  "Jenkins",
				  "Docker",
				  "JIRA"
				],
				"description": "Intermediate"
			  }
			],
			"columns": 1,
			"visible": true
		  },
		  "summary": {
			"id": "summary",
			"name": "Summary",
			"columns": 1,
			"content": "<p>Innovative Web Developer with 5 years of experience in building impactful and user-friendly websites and applications. Specializes in <strong>front-end technologies</strong> and passionate about modern web standards and cutting-edge development techniques. Proven track record of leading successful projects from concept to deployment.</p>",
			"visible": true
		  },
		  "profiles": {
			"id": "profiles",
			"name": "Profiles",
			"items": [
			  {
				"id": "cnbk5f0aeqvhx69ebk7hktwd",
				"url": {
				  "href": "https://linkedin.com/in/johndoe",
				  "label": ""
				},
				"icon": "linkedin",
				"network": "LinkedIn",
				"visible": true,
				"username": "johndoe"
			  },
			  {
				"id": "ukl0uecvzkgm27mlye0wazlb",
				"url": {
				  "href": "https://github.com/johndoe",
				  "label": ""
				},
				"icon": "github",
				"network": "GitHub",
				"visible": true,
				"username": "johndoe"
			  }
			],
			"columns": 1,
			"visible": true
		  },
		  "projects": {
			"id": "projects",
			"name": "Projects",
			"items": [
			  {
				"id": "yw843emozcth8s1ubi1ubvlf",
				"url": {
				  "href": "",
				  "label": ""
				},
				"date": "",
				"name": "E-Commerce Platform",
				"summary": "<p>Led the development of a full-stack e-commerce platform, improving sales conversion by 25%.</p>",
				"visible": true,
				"keywords": [],
				"description": "Project Lead"
			  },
			  {
				"id": "ncxgdjjky54gh59iz2t1xi1v",
				"url": {
				  "href": "",
				  "label": ""
				},
				"date": "",
				"name": "Interactive Dashboard",
				"summary": "<p>Created an interactive analytics dashboard for a SaaS application, enhancing data visualization for clients.</p>",
				"visible": true,
				"keywords": [],
				"description": "Frontend Developer"
			  }
			],
			"columns": 1,
			"visible": true
		  },
		  "education": {
			"id": "education",
			"name": "Education",
			"items": [
			  {
				"id": "yo3p200zo45c6cdqc6a2vtt3",
				"url": {
				  "href": "",
				  "label": ""
				},
				"area": "Berkeley, CA",
				"date": "August 2012 to May 2016",
				"score": "",
				"summary": "",
				"visible": true,
				"studyType": "Bachelor's in Computer Science",
				"institution": "University of California"
			  }
			],
			"columns": 1,
			"visible": true
		  },
		  "interests": {
			"id": "interests",
			"name": "Interests",
			"items": [],
			"columns": 1,
			"visible": true
		  },
		  "languages": {
			"id": "languages",
			"name": "Languages",
			"items": [],
			"columns": 1,
			"visible": true
		  },
		  "volunteer": {
			"id": "volunteer",
			"name": "Volunteering",
			"items": [],
			"columns": 1,
			"visible": true
		  },
		  "experience": {
			"id": "experience",
			"name": "Experience",
			"items": [
			  {
				"id": "lhw25d7gf32wgdfpsktf6e0x",
				"url": {
				  "href": "https://creativesolutions.inc/",
				  "label": ""
				},
				"date": "January 2019 to Present",
				"company": "Creative Solutions Inc.",
				"summary": "<ul><li><p>Spearheaded the redesign of the main product website, resulting in a 40% increase in user engagement.</p></li><li><p>Developed and implemented a new responsive framework, improving cross-device compatibility.</p></li><li><p>Mentored a team of four junior developers, fostering a culture of technical excellence.</p></li></ul>",
				"visible": true,
				"location": "San Francisco, CA",
				"position": "Senior Web Developer"
			  },
			  {
				"id": "r6543lil53ntrxmvel53gbtm",
				"url": {
				  "href": "https://techadvancers.com/",
				  "label": ""
				},
				"date": "June 2016 to December 2018",
				"company": "TechAdvancers",
				"summary": "<ul><li><p>Collaborated in a team of 10 to develop high-quality web applications using React.js and Node.js.</p></li><li><p>Managed the integration of third-party services such as Stripe for payments and Twilio for SMS services.</p></li><li><p>Optimized application performance, achieving a 30% reduction in load times.</p></li></ul>",
				"visible": true,
				"location": "San Jose, CA",
				"position": "Web Developer"
			  }
			],
			"columns": 1,
			"visible": true
		  },
		  "references": {
			"id": "references",
			"name": "References",
			"items": [
			  {
				"id": "f2sv5z0cce6ztjl87yuk8fak",
				"url": {
				  "href": "",
				  "label": ""
				},
				"name": "Available upon request",
				"summary": "",
				"visible": true,
				"description": ""
			  }
			],
			"columns": 1,
			"visible": false
		  },
		  "publications": {
			"id": "publications",
			"name": "Publications",
			"items": [],
			"columns": 1,
			"visible": true
		  },
		  "certifications": {
			"id": "certifications",
			"name": "Certifications",
			"items": [
			  {
				"id": "spdhh9rrqi1gvj0yqnbqunlo",
				"url": {
				  "href": "",
				  "label": ""
				},
				"date": "2020",
				"name": "Full-Stack Web Development",
				"issuer": "CodeAcademy",
				"summary": "",
				"visible": true
			  },
			  {
				"id": "n838rddyqv47zexn6cxauwqp",
				"url": {
				  "href": "",
				  "label": ""
				},
				"date": "2019",
				"name": "AWS Certified Developer",
				"issuer": "Amazon Web Services",
				"summary": "",
				"visible": true
			  }
			],
			"columns": 1,
			"visible": true
		  }
		}
	  }`)
	rsm, err := getResumeFromJSON(data)
	if err != nil {
		t.Errorf("Error importing JSON: %s", err)
	}
	want := "John Doe"
	if got := rsm.Basics.Name; got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}

	t.Run("Save as JSON", func(t *testing.T) {
		_, err := rsm.toJSON()
		if err != nil {
			t.Errorf("Got an error: %v", err)
		}
	})
}

// func TestGetAndExportResume(t *testing.T) {
// 	// Create file for testing
// 	fname1 := "test in .json"
// 	fname2 := "test out.json"
// 	f, err := os.Create(fname1)
// 	defer os.Remove(fname1)
// 	if err != nil {
// 		t.Errorf("Error creating test file: %s", err)
// 	}
// 	f.Write([]byte("{\n\t\t\"basics\": {\n\t\t  \"url\": {\n\t\t\t\"href\": \"https://johndoe.me/\",\n\t\t\t\"label\": \"\"\n\t\t  },\n\t\t  \"name\": \"John Doe\",\n\t\t  \"email\": \"john.doe@gmail.com\",\n\t\t  \"phone\": \"(555) 123-4567\",\n\t\t  \"picture\": {\n\t\t\t\"url\": \"https://i.imgur.com/HgwyOuJ.jpg\",\n\t\t\t\"size\": 120,\n\t\t\t\"effects\": {\n\t\t\t  \"border\": false,\n\t\t\t  \"hidden\": false,\n\t\t\t  \"grayscale\": false\n\t\t\t},\n\t\t\t\"aspectRatio\": 1,\n\t\t\t\"borderRadius\": 0\n\t\t  },\n\t\t  \"headline\": \"Creative and Innovative Web Developer\",\n\t\t  \"location\": \"Pleasantville, CA 94588\",\n\t\t  \"customFields\": []\n\t\t},\n\t\t\"metadata\": {\n\t\t  \"css\": {\n\t\t\t\"value\": \".section {\\n\\toutline: 1px solid #000;\\n\\toutline-offset: 4px;\\n}\",\n\t\t\t\"visible\": false\n\t\t  },\n\t\t  \"page\": {\n\t\t\t\"format\": \"a4\",\n\t\t\t\"margin\": 14,\n\t\t\t\"options\": {\n\t\t\t  \"breakLine\": true,\n\t\t\t  \"pageNumbers\": true\n\t\t\t}\n\t\t  },\n\t\t  \"notes\": \"\",\n\t\t  \"theme\": {\n\t\t\t\"text\": \"#000000\",\n\t\t\t\"primary\": \"#ca8a04\",\n\t\t\t\"background\": \"#ffffff\"\n\t\t  },\n\t\t  \"layout\": [\n\t\t\t[\n\t\t\t  [\n\t\t\t\t\"summary\",\n\t\t\t\t\"experience\",\n\t\t\t\t\"education\",\n\t\t\t\t\"projects\",\n\t\t\t\t\"references\"\n\t\t\t  ],\n\t\t\t  [\n\t\t\t\t\"profiles\",\n\t\t\t\t\"skills\",\n\t\t\t\t\"certifications\",\n\t\t\t\t\"interests\",\n\t\t\t\t\"languages\",\n\t\t\t\t\"awards\",\n\t\t\t\t\"volunteer\",\n\t\t\t\t\"publications\"\n\t\t\t  ]\n\t\t\t]\n\t\t  ],\n\t\t  \"template\": \"glalie\",\n\t\t  \"typography\": {\n\t\t\t\"font\": {\n\t\t\t  \"size\": 13,\n\t\t\t  \"family\": \"Merriweather\",\n\t\t\t  \"subset\": \"latin\",\n\t\t\t  \"variants\": [\n\t\t\t\t\"regular\"\n\t\t\t  ]\n\t\t\t},\n\t\t\t\"hideIcons\": false,\n\t\t\t\"lineHeight\": 1.75,\n\t\t\t\"underlineLinks\": true\n\t\t  }\n\t\t},\n\t\t\"sections\": {\n\t\t  \"awards\": {\n\t\t\t\"id\": \"awards\",\n\t\t\t\"name\": \"Awards\",\n\t\t\t\"items\": [],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"custom\": {},\n\t\t  \"skills\": {\n\t\t\t\"id\": \"skills\",\n\t\t\t\"name\": \"Skills\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"hn0keriukh6c0ojktl9gsgjm\",\n\t\t\t\t\"name\": \"Web Technologies\",\n\t\t\t\t\"level\": 0,\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"keywords\": [\n\t\t\t\t  \"HTML5\",\n\t\t\t\t  \"JavaScript\",\n\t\t\t\t  \"PHP\",\n\t\t\t\t  \"Python\"\n\t\t\t\t],\n\t\t\t\t\"description\": \"Advanced\"\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"r8c3y47vykausqrgmzwg5pur\",\n\t\t\t\t\"name\": \"Web Frameworks\",\n\t\t\t\t\"level\": 0,\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"keywords\": [\n\t\t\t\t  \"React.js\",\n\t\t\t\t  \"Angular\",\n\t\t\t\t  \"Vue.js\",\n\t\t\t\t  \"Laravel\",\n\t\t\t\t  \"Django\"\n\t\t\t\t],\n\t\t\t\t\"description\": \"Intermediate\"\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"b5l75aseexqv17quvqgh73fe\",\n\t\t\t\t\"name\": \"Tools\",\n\t\t\t\t\"level\": 0,\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"keywords\": [\n\t\t\t\t  \"Webpack\",\n\t\t\t\t  \"Git\",\n\t\t\t\t  \"Jenkins\",\n\t\t\t\t  \"Docker\",\n\t\t\t\t  \"JIRA\"\n\t\t\t\t],\n\t\t\t\t\"description\": \"Intermediate\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"summary\": {\n\t\t\t\"id\": \"summary\",\n\t\t\t\"name\": \"Summary\",\n\t\t\t\"columns\": 1,\n\t\t\t\"content\": \"<p>Innovative Web Developer with 5 years of experience in building impactful and user-friendly websites and applications. Specializes in <strong>front-end technologies</strong> and passionate about modern web standards and cutting-edge development techniques. Proven track record of leading successful projects from concept to deployment.</p>\",\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"profiles\": {\n\t\t\t\"id\": \"profiles\",\n\t\t\t\"name\": \"Profiles\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"cnbk5f0aeqvhx69ebk7hktwd\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"https://linkedin.com/in/johndoe\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"icon\": \"linkedin\",\n\t\t\t\t\"network\": \"LinkedIn\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"username\": \"johndoe\"\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"ukl0uecvzkgm27mlye0wazlb\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"https://github.com/johndoe\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"icon\": \"github\",\n\t\t\t\t\"network\": \"GitHub\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"username\": \"johndoe\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"projects\": {\n\t\t\t\"id\": \"projects\",\n\t\t\t\"name\": \"Projects\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"yw843emozcth8s1ubi1ubvlf\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"\",\n\t\t\t\t\"name\": \"E-Commerce Platform\",\n\t\t\t\t\"summary\": \"<p>Led the development of a full-stack e-commerce platform, improving sales conversion by 25%.</p>\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"keywords\": [],\n\t\t\t\t\"description\": \"Project Lead\"\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"ncxgdjjky54gh59iz2t1xi1v\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"\",\n\t\t\t\t\"name\": \"Interactive Dashboard\",\n\t\t\t\t\"summary\": \"<p>Created an interactive analytics dashboard for a SaaS application, enhancing data visualization for clients.</p>\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"keywords\": [],\n\t\t\t\t\"description\": \"Frontend Developer\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"education\": {\n\t\t\t\"id\": \"education\",\n\t\t\t\"name\": \"Education\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"yo3p200zo45c6cdqc6a2vtt3\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"area\": \"Berkeley, CA\",\n\t\t\t\t\"date\": \"August 2012 to May 2016\",\n\t\t\t\t\"score\": \"\",\n\t\t\t\t\"summary\": \"\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"studyType\": \"Bachelor's in Computer Science\",\n\t\t\t\t\"institution\": \"University of California\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"interests\": {\n\t\t\t\"id\": \"interests\",\n\t\t\t\"name\": \"Interests\",\n\t\t\t\"items\": [],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"languages\": {\n\t\t\t\"id\": \"languages\",\n\t\t\t\"name\": \"Languages\",\n\t\t\t\"items\": [],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"volunteer\": {\n\t\t\t\"id\": \"volunteer\",\n\t\t\t\"name\": \"Volunteering\",\n\t\t\t\"items\": [],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"experience\": {\n\t\t\t\"id\": \"experience\",\n\t\t\t\"name\": \"Experience\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"lhw25d7gf32wgdfpsktf6e0x\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"https://creativesolutions.inc/\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"January 2019 to Present\",\n\t\t\t\t\"company\": \"Creative Solutions Inc.\",\n\t\t\t\t\"summary\": \"<ul><li><p>Spearheaded the redesign of the main product website, resulting in a 40% increase in user engagement.</p></li><li><p>Developed and implemented a new responsive framework, improving cross-device compatibility.</p></li><li><p>Mentored a team of four junior developers, fostering a culture of technical excellence.</p></li></ul>\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"location\": \"San Francisco, CA\",\n\t\t\t\t\"position\": \"Senior Web Developer\"\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"r6543lil53ntrxmvel53gbtm\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"https://techadvancers.com/\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"June 2016 to December 2018\",\n\t\t\t\t\"company\": \"TechAdvancers\",\n\t\t\t\t\"summary\": \"<ul><li><p>Collaborated in a team of 10 to develop high-quality web applications using React.js and Node.js.</p></li><li><p>Managed the integration of third-party services such as Stripe for payments and Twilio for SMS services.</p></li><li><p>Optimized application performance, achieving a 30% reduction in load times.</p></li></ul>\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"location\": \"San Jose, CA\",\n\t\t\t\t\"position\": \"Web Developer\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"references\": {\n\t\t\t\"id\": \"references\",\n\t\t\t\"name\": \"References\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"f2sv5z0cce6ztjl87yuk8fak\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"name\": \"Available upon request\",\n\t\t\t\t\"summary\": \"\",\n\t\t\t\t\"visible\": true,\n\t\t\t\t\"description\": \"\"\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": false\n\t\t  },\n\t\t  \"publications\": {\n\t\t\t\"id\": \"publications\",\n\t\t\t\"name\": \"Publications\",\n\t\t\t\"items\": [],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  },\n\t\t  \"certifications\": {\n\t\t\t\"id\": \"certifications\",\n\t\t\t\"name\": \"Certifications\",\n\t\t\t\"items\": [\n\t\t\t  {\n\t\t\t\t\"id\": \"spdhh9rrqi1gvj0yqnbqunlo\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"2020\",\n\t\t\t\t\"name\": \"Full-Stack Web Development\",\n\t\t\t\t\"issuer\": \"CodeAcademy\",\n\t\t\t\t\"summary\": \"\",\n\t\t\t\t\"visible\": true\n\t\t\t  },\n\t\t\t  {\n\t\t\t\t\"id\": \"n838rddyqv47zexn6cxauwqp\",\n\t\t\t\t\"url\": {\n\t\t\t\t  \"href\": \"\",\n\t\t\t\t  \"label\": \"\"\n\t\t\t\t},\n\t\t\t\t\"date\": \"2019\",\n\t\t\t\t\"name\": \"AWS Certified Developer\",\n\t\t\t\t\"issuer\": \"Amazon Web Services\",\n\t\t\t\t\"summary\": \"\",\n\t\t\t\t\"visible\": true\n\t\t\t  }\n\t\t\t],\n\t\t\t\"columns\": 1,\n\t\t\t\"visible\": true\n\t\t  }\n\t\t}\n\t  }"))
// 	f.Close()

// 	rsm, err := GetResume(fname1)
// 	if err != nil {
// 		t.Errorf("Error importing JSON: %s", err)
// 	}
// 	want := "John Doe"
// 	if got := rsm.Basics.Name; got != want {
// 		t.Errorf("Got: %v, Want: %v", got, want)
// 	}

// 	// Translate
// 	rsmTranslated, err := rsm.Translate("YOUR DEEPL API AUTH KEY", "NL")
// 	if err != nil {
// 		t.Errorf("Error translating resume: %s", err)
// 	}
// 	want = "Creatieve en innovatieve webontwikkelaar"
// 	if got := rsmTranslated.Basics.Headline; got != want {
// 		t.Errorf("Got: %v, Want: %v", got, want)
// 	}

// 	err = rsmTranslated.Export(fname2)
// 	//defer os.Remove(fname2)
// 	if err != nil {
// 		t.Errorf("Error exporting JSON: %s", err)
// 	}

// }
