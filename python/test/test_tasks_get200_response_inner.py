# coding: utf-8

"""
    app.salesflare.com

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from salesflare.models.tasks_get200_response_inner import TasksGet200ResponseInner

class TestTasksGet200ResponseInner(unittest.TestCase):
    """TasksGet200ResponseInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> TasksGet200ResponseInner:
        """Test TasksGet200ResponseInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `TasksGet200ResponseInner`
        """
        model = TasksGet200ResponseInner()
        if include_optional:
            return TasksGet200ResponseInner(
                id = 56,
                type = '',
                account = salesflare.models._tasks_get_200_response_inner_account._tasks_get_200_response_inner_account(
                    id = 56, 
                    name = '', 
                    picture = '', ),
                creator = 56,
                description = '',
                reminder_date = '',
                email = None,
                company = None,
                meeting = None,
                completed = True,
                completion_date = None,
                completor = None,
                archived = True,
                archive_date = None,
                archivor = None,
                last_modified_by = 56,
                creation_date = '',
                modification_date = '',
                assignees = [
                    salesflare.models._tasks_get_200_response_inner_assignees_inner._tasks_get_200_response_inner_assignees_inner(
                        id = 56, 
                        email = '', 
                        picture = '', 
                        name = '', 
                        disabled = True, )
                    ],
                can_edit = True
            )
        else:
            return TasksGet200ResponseInner(
                id = 56,
                type = '',
                account = salesflare.models._tasks_get_200_response_inner_account._tasks_get_200_response_inner_account(
                    id = 56, 
                    name = '', 
                    picture = '', ),
                creator = 56,
                description = '',
                reminder_date = '',
                email = None,
                company = None,
                meeting = None,
                completed = True,
                completion_date = None,
                completor = None,
                archived = True,
                archive_date = None,
                archivor = None,
                last_modified_by = 56,
                creation_date = '',
                modification_date = '',
                assignees = [
                    salesflare.models._tasks_get_200_response_inner_assignees_inner._tasks_get_200_response_inner_assignees_inner(
                        id = 56, 
                        email = '', 
                        picture = '', 
                        name = '', 
                        disabled = True, )
                    ],
                can_edit = True,
        )
        """

    def testTasksGet200ResponseInner(self):
        """Test TasksGet200ResponseInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
