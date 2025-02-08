SELECT * FROM public.episodes
ORDER BY id ASC 

SELECT id, phrase,you_tube_url, embedding FROM episodes;
SELECT id, phrase,you_tube_url, embedding FROM episodes order by 1 desc;



INSERT INTO public.episodes(
	id, created_at, phrase, episode, title, "timestamp", you_tube_url)
	VALUES (106, current_timestamp, 'Cuidado cuidado... te pasaste', 'S06E01', 'Bart of Darkness', '0:09', 'https://www.youtube.com/watch?v=den_9jIKwfk');

