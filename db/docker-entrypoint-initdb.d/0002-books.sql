COPY public.books (id, title, author, isbn, year, publisher, tags, max_loan_days, location, format) FROM stdin;
1	Jane Eyre	Charlotte Brontë	978-0141441146	1847	Penguin Classics	{classics,reprints}	21	Classic Fiction	Paperback
2	Wuthering Heights	Emily Jane Brontë	0-486-29256-8	1847	Thomas Cautley Newby	{classics,"rare books"}	0	Rare Books Collection	Hardcover
3	Agnes Grey	Anne Brontë	978014043210	1850	Pengin Classics	{classics,reprints}	21	Classic Fiction	Hardcover
4	Jane Eyre	Charlotte Brontë	978-0141441146	1847	Penguin Classics	{classics,reprints,english,"19th century"}	21	Classic Fiction	Paperback
5	Wuthering Heights	Emily Jane Brontë	0-486-29256-8	1847	Thomas Cautley Newby	{classics,"rare books",english,"19th century"}	0	Rare Books Collection	Hardcover
6	Agnes Grey	Anne Brontë	978014043210	1850	Pengin Classics	{classics,reprints,english,"19th century"}	21	Classic Fiction	Hardcover
7	My Brilliant Career	Miles Franklin	9781600963780	1901	Waking Lion Press	{classics,australian,reprints,"20th century"}	21	Australian Classics	Paperback
8	Cloudstreet	Tim Winton	9780743234412	1991	Scribner	{australian,reprints,"20th century"}	21	Australian Literature	Paperback
9	Picnic at Hanging Rock	Joan Lindsay	n/a	1967	n/a	{australian,classic,"rare books",manuscript,"20th century"}	0	Original Manuscripts Collection	Manuscript
10	For the Term of His Natural Life	Marcus Clarke	9781406512038	1874	Dodo Press	{australian,classic,"19th century"}	21	Australian Classics	Paperback
11	The Secret River	Kate Grenville	9781841959146	2005	Canongate U.S.	{australian,"21st century"}	21	Australian Literature	Hardcover
\.
