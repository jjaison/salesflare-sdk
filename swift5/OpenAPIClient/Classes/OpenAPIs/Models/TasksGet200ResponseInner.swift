//
// TasksGet200ResponseInner.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

public struct TasksGet200ResponseInner: Codable, JSONEncodable, Hashable {

    public var id: Int
    public var type: String
    public var account: TasksGet200ResponseInnerAccount
    public var creator: Int?
    public var description: String
    public var reminderDate: String
    public var email: TasksGet200ResponseInnerEmail
    public var company: AnyCodable?
    public var meeting: AnyCodable?
    public var completed: Bool
    public var completionDate: AnyCodable?
    public var completor: AnyCodable?
    public var archived: Bool
    public var archiveDate: AnyCodable?
    public var archivor: AnyCodable?
    public var lastModifiedBy: Int?
    public var creationDate: String
    public var modificationDate: String
    public var assignees: [TasksGet200ResponseInnerAssigneesInner]
    public var canEdit: Bool

    public init(id: Int, type: String, account: TasksGet200ResponseInnerAccount, creator: Int?, description: String, reminderDate: String, email: TasksGet200ResponseInnerEmail, company: AnyCodable?, meeting: AnyCodable?, completed: Bool, completionDate: AnyCodable?, completor: AnyCodable?, archived: Bool, archiveDate: AnyCodable?, archivor: AnyCodable?, lastModifiedBy: Int?, creationDate: String, modificationDate: String, assignees: [TasksGet200ResponseInnerAssigneesInner], canEdit: Bool) {
        self.id = id
        self.type = type
        self.account = account
        self.creator = creator
        self.description = description
        self.reminderDate = reminderDate
        self.email = email
        self.company = company
        self.meeting = meeting
        self.completed = completed
        self.completionDate = completionDate
        self.completor = completor
        self.archived = archived
        self.archiveDate = archiveDate
        self.archivor = archivor
        self.lastModifiedBy = lastModifiedBy
        self.creationDate = creationDate
        self.modificationDate = modificationDate
        self.assignees = assignees
        self.canEdit = canEdit
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case id
        case type
        case account
        case creator
        case description
        case reminderDate = "reminder_date"
        case email
        case company
        case meeting
        case completed
        case completionDate = "completion_date"
        case completor
        case archived
        case archiveDate = "archive_date"
        case archivor
        case lastModifiedBy = "last_modified_by"
        case creationDate = "creation_date"
        case modificationDate = "modification_date"
        case assignees
        case canEdit = "can_edit"
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(id, forKey: .id)
        try container.encode(type, forKey: .type)
        try container.encode(account, forKey: .account)
        try container.encode(creator, forKey: .creator)
        try container.encode(description, forKey: .description)
        try container.encode(reminderDate, forKey: .reminderDate)
        try container.encode(email, forKey: .email)
        try container.encode(company, forKey: .company)
        try container.encode(meeting, forKey: .meeting)
        try container.encode(completed, forKey: .completed)
        try container.encode(completionDate, forKey: .completionDate)
        try container.encode(completor, forKey: .completor)
        try container.encode(archived, forKey: .archived)
        try container.encode(archiveDate, forKey: .archiveDate)
        try container.encode(archivor, forKey: .archivor)
        try container.encode(lastModifiedBy, forKey: .lastModifiedBy)
        try container.encode(creationDate, forKey: .creationDate)
        try container.encode(modificationDate, forKey: .modificationDate)
        try container.encode(assignees, forKey: .assignees)
        try container.encode(canEdit, forKey: .canEdit)
    }
}

