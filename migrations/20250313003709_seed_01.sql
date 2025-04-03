-- +goose Up
-- +goose StatementBegin
INSERT INTO environments VALUES(1287647232,'HML','description','hml, dev, beholder','http://localhost:8080',CURRENT_TIMESTAMP,NULL);
INSERT INTO sessions VALUES(511709184,1287647232,'user1','tag1, tag2',CURRENT_TIMESTAMP,NULL);
INSERT INTO sessions VALUES(511709187,null,'user1','tag1, tag2',CURRENT_TIMESTAMP,NULL);
INSERT INTO requests VALUES(469766144,1287647232,511709184,'create foo 3','','POST','/foo','{''foo'' : ''bar edited againnn''}',NULL,'2006-01-04T15:10:05Z','2025-03-13T00:26:39.600498306Z',NULL);
INSERT INTO requests VALUES(901779456,1287647232,511709187,'update project language','','PUT','/1234/languages/6789','{ ''name'' : ''Português'' }',NULL,'2025-03-11T02:10:05Z','2025-03-13T00:32:12.503485173Z',NULL);
INSERT INTO requests VALUES(1056968704,1287647232,511709187,'delete project language','','DELETE','/1234/languages/6789',NULL,NULL,'2025-03-11T05:18:05Z','2025-03-13T00:32:35.068372292Z',NULL);
INSERT INTO requests VALUES(1103106048,1287647232,511709184,'create foo ','','POST','/foo','{''foo'' : ''bar edited againnn''}',NULL,'2006-01-02T15:10:05Z','2025-03-13T00:26:30.535479043Z',NULL);
INSERT INTO requests VALUES(1526730752,1287647232,511709184,'create foo 2','','POST','/foo','{''foo'' : ''bar edited againnn''}',NULL,'2006-01-03T15:10:05Z','2025-03-13T00:26:34.732304568Z',NULL);
INSERT INTO requests VALUES(1971326976,1287647232,511709187,'create project language','','POST','/1234/languages','{''language'' : {''name'' : ''Portuguese'', ''slug'' : ''pt-BR''} }',NULL,'2025-03-11T00:10:05Z','2025-03-13T00:31:26.678592881Z',NULL);
INSERT INTO requests VALUES(2004881408,1287647232,511709187,'get project languages','','GET','/1234/languages',NULL,NULL,'2025-01-04T15:10:05Z','2025-03-13T00:30:11.934229183Z',NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM requests;
-- +goose StatementEnd
