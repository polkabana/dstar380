package main

const codes string = `289	Abkhazia	GE-AB
412	Afghanistan	AF
276	Albania	AL
603	Algeria	DZ
544	American Samoa (United States of America)	AS
213	Andorra	AD
631	Angola	AO
365	Anguilla (United Kingdom)	AI
344	Antigua and Barbuda	AG
722	Argentina	AR
283	Armenia	AM
363	Aruba	AW
505	Australia	AU
232	Austria	AT
400	Azerbaijan	AZ
364	Bahamas	BS
426	Bahrain	BH
470	Bangladesh	BD
342	Barbados	BB
257	Belarus	BY
206	Belgium	BE
702	Belize	BZ
616	Benin	BJ
350	Bermuda	BM
402	Bhutan	BT
736	Bolivia	BO
362	Bonaire, Sint Eustatius and Saba	BQ
218	Bosnia and Herzegovina	BA
652	Botswana	BW
724	Brazil	BR
995	British Indian Ocean Territory (United Kingdom)	IO
348	British Virgin Islands (United Kingdom)	VG
528	Brunei	BN
284	Bulgaria	BG
613	Burkina Faso	BF
642	Burundi	BI
456	Cambodia	KH
624	Cameroon	CM
302	Canada	CA
625	Cape Verde	CV
346	Cayman Islands (United Kingdom)	KY
623	Central African Republic	CF
622	Chad	TD
730	Chile	CL
460	China	CN
461	China	CN
732	Colombia	CO
654	Comoros	KM
629	Congo	CG
548	Cook Islands (Pacific Ocean)	CK
712	Costa Rica	CR
219	Croatia	HR
368	Cuba	CU
362	Curaçao	CW
280	Cyprus	CY
230	Czech Republic	CZ
630	Democratic Republic of the Congo	CD
238	Denmark (Kingdom of Denmark)	DK
638	Djibouti	DJ
366	Dominica	DM
370	Dominican Republic	DO
514	East Timor	TL
740	Ecuador	EC
602	Egypt	EG
706	El Salvador	SV
627	Equatorial Guinea	GQ
657	Eritrea	ER
248	Estonia	EE
636	Ethiopia	ET
750	Falkland Islands (United Kingdom)	FK
288	Faroe Islands (Kingdom of Denmark)	FO
542	Fiji	FJ
244	Finland	FI
208	France	FR
742	French Guiana (France)	GF
647	French Indian Ocean Territories (France)	RE
647	French Indian Ocean Territories (France)	YT
547	French Polynesia (France)	PF
628	Gabon	GA
607	Gambia	GM
282	Georgia	GE
262	Germany	DE
620	Ghana	GH
266	Gibraltar (United Kingdom)	GI
202	Greece	GR
290	Greenland (Kingdom of Denmark)	GL
352	Grenada	GD
340	Guadeloupe (France)	GP
310	Guam (United States of America)	GU
311	Guam (United States of America)	GU
704	Guatemala	GT
234	Guernsey (United Kingdom)	GG
611	Guinea	GN
632	Guinea-Bissau	GW
738	Guyana	GY
372	Haiti	HT
708	Honduras	HN
454	Hong Kong	HK
216	Hungary	HU
274	Iceland	IS
404	India	IN
405	India	IN
406	India	IN
510	Indonesia	ID
432	Iran	IR
418	Iraq	IQ
272	Ireland	IE
234	Isle of Man (United Kingdom)	IM
425	Israel	IL
222	Italy	IT
612	Ivory Coast	CI
338	Jamaica	JM
440	Japan	JP
441	Japan	JP
234	Jersey (United Kingdom)	JE
416	Jordan	JO
401	Kazakhstan	KZ
639	Kenya	KE
545	Kiribati	KI
467	Korea, North	KP
450	Korea, South	KR
221	Kosovo	XK
419	Kuwait	KW
437	Kyrgyzstan	KG
457	Laos	LA
247	Latvia	LV
415	Lebanon	LB
651	Lesotho	LS
618	Liberia	LR
606	Libya	LY
295	Liechtenstein	LI
246	Lithuania	LT
270	Luxembourg	LU
455	Macau (People's Republic of China)	MO
646	Madagascar	MG
650	Malawi	MW
502	Malaysia	MY
472	Maldives	MV
610	Mali	ML
278	Malta	MT
551	Marshall Islands	MH
340	Martinique (France)	MQ
609	Mauritania	MR
617	Mauritius	MU
334	Mexico	MX
550	Micronesia, Federated States of	FM
259	Moldova	MD
212	Monaco	MC
428	Mongolia	MN
297	Montenegro	ME
354	Montserrat (United Kingdom)	MS
604	Morocco	MA
643	Mozambique	MZ
414	Myanmar	MM
649	Namibia	NA
536	Nauru	NR
429	Nepal	NP
204	Netherlands (Kingdom of the Netherlands)	NL
546	New Caledonia	NC
530	New Zealand	NZ
710	Nicaragua	NI
614	Niger	NE
621	Nigeria	NG
555	Niue	NU
505	Norfolk Island	NF
294	North Macedonia	MK
310	Northern Mariana Islands (United States of America)	MP
242	Norway	NO
422	Oman	OM
410	Pakistan	PK
552	Palau	PW
425	Palestine	PS
714	Panama	PA
537	Papua New Guinea	PG
744	Paraguay	PY
716	Peru	PE
515	Philippines	PH
260	Poland	PL
268	Portugal	PT
330	Puerto Rico	PR
427	Qatar	QA
226	Romania	RO
250	Russian Federation	RU
635	Rwanda	RW
340	Saint Barthélemy (France)	BL
658	Saint Helena, Ascension and Tristan da Cunha	SH
356	Saint Kitts and Nevis	KN
358	Saint Lucia	LC
340	Saint Martin (France)	MF
308	Saint Pierre and Miquelon	PM
360	Saint Vincent and the Grenadines	VC
549	Samoa	WS
292	San Marino	SM
626	Sao Tome and Principe	ST
420	Saudi Arabia	SA
608	Senegal	SN
220	Serbia	RS
633	Seychelles	SC
619	Sierra Leone	SL
525	Singapore	SG
362	Sint Maarten	SX
231	Slovakia	SK
293	Slovenia	SI
540	Solomon Islands	SB
637	Somalia	SO
655	South Africa	ZA
659	South Sudan	SS
214	Spain	ES
413	Sri Lanka	LK
634	Sudan	SD
746	Suriname	SR
653	Swaziland	SZ
240	Sweden	SE
228	Switzerland	CH
417	Syria	SY
466	Taiwan	TW
436	Tajikistan	TJ
640	Tanzania	TZ
520	Thailand	TH
615	Togo	TG
554	Tokelau	TK
539	Tonga	TO
374	Trinidad and Tobago	TT
605	Tunisia	TN
286	Turkey	TR
438	Turkmenistan	TM
376	Turks and Caicos Islands	TC
553	Tuvalu	TV
641	Uganda	UG
255	Ukraine	UA
424	United Arab Emirates	AE
430	United Arab Emirates (Abu Dhabi)	AE
431	United Arab Emirates (Dubai)	AE
234	United Kingdom	GB
235	United Kingdom	GB
310	United States of America	US
311	United States of America	US
312	United States of America	US
313	United States of America	US
314	United States of America	US
315	United States of America	US
316	United States of America	US
317	United States of America	US
320	United States of America	US
332	United States Virgin Islands	VI
748	Uruguay	UY
434	Uzbekistan	UZ
541	Vanuatu	VU
734	Venezuela	VE
452	Vietnam	VN
543	Wallis and Futuna	WF
421	Yemen	YE
645	Zambia	ZM
648	Zimbabwe	ZW
`
