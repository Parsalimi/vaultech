# Functional Requirements Specification (FRS)

## CSVault / Vaultech

**Version:** MVP v0.1 (Authentication Decisions Incorporated)  
**Document Type:** Functional Requirements Specification  
**Status:** Draft / Requirements Baseline

---

# 1. Document Purpose

This document defines the functional requirements for the CSVault / Vaultech website MVP v0.1.

The purpose of this FRS is to describe **what the system shall do**, the actors who interact with it, the main domain concepts, access rules, workflows, and functional constraints.

This document is intentionally organized by **domain and capability**, rather than by the order in which requirements were originally written.

---

# 2. Scope

The MVP covers the following major capabilities:

1. Authentication and user management
2. Academic term management
3. Course and course-offering management
4. Course content and resources
5. Course calendar and event management
6. Professor discovery ("Ostad Shenasi")
7. Professor profiles
8. Professor reviews
9. Review eligibility and duplicate prevention
10. Review moderation and reporting
11. Professor evaluation criteria and features
12. Administrative management
13. Rating aggregation and review lifecycle
14. Privacy and historical-data preservation

---

# 3. Actors and Roles

## 3.1 Student

A Student can:

- Authenticate and maintain a student profile.
- Search and browse courses.
- Add and remove courses from My Courses.
- View course content according to permissions.
- Browse courses by academic term.
- Discover professors.
- View professor profiles and published reviews.
- Submit eligible professor reviews.
- Report published reviews.

## 3.2 Teaching Assistant (TA)

A TA is a Student who has been assigned as a Teaching Assistant for a specific Course.

A TA:

- Retains all normal Student capabilities.
- Receives additional permissions for assigned courses.
- Does **not** become a separate global user role from Student.
- Cannot modify administrator-controlled academic information unless explicitly authorized.

## 3.3 Professor

A Professor can:

- Access courses they teach.
- Manage permitted course information and resources.
- Manage announcements, materials, assignments, exams, and calendar-related information.
- Manage actual course Meeting Times.
- Cannot modify administrator-controlled academic information.

## 3.4 Administrator

An Administrator can:

- Manage users and professors.
- Manage academic terms.
- Manage courses and course offerings.
- Assign instructors and TAs.
- Manage administrator-controlled course information.
- Moderate reviews and review reports.
- Manage evaluation criteria and professor features.
- Perform administrative overrides where specified.

---

# 4. Domain Model and Terminology

## 4.1 Academic Term

An Academic Term represents a specific academic period during which Course Offerings are conducted.

Each Academic Term contains:

- Term name / season
- Academic year
- Academic Term Code

Example:

```text
4051 — زمستان 1405
```

The exact naming and coding convention shall be defined by the platform's academic calendar configuration.

## 4.2 Course

A Course represents the persistent academic definition of a course.

Examples of Course-level information include:

- Course Code
- Course Name
- Degree Level
- Major
- Department
- Year of Entry
- Type
- Prerequisites
- Description

## 4.3 Course Offering

A Course Offering represents a specific occurrence of a Course in an Academic Term.

A Course Offering may contain:

- Academic Term
- Instructor
- TAs
- Official Presentation Time
- Meeting Times
- Syllabus
- Objectives
- References / Textbooks
- Grading Policy
- Announcements
- Materials
- Assignments
- Exams
- Calendar events

A Course Offering is therefore the academic-context-specific instance of a Course.

## 4.4 Official Presentation Time

Presentation Time is the official time at which the university has scheduled the course.

It is administrative academic information.

- Administrators may manage it.
- Professors and TAs shall not modify it.
- Actual schedule changes shall be represented through Meeting Times.

## 4.5 Meeting Time

A Meeting Time represents an actual scheduled session associated with a Course Offering.

It may contain:

- Day
- Start Time
- End Time
- Location
- Meeting Type

Possible Meeting Types include:

- Regular Lecture
- Exercise / Class
- Office Hour
- Q&A Session
- Make-up Class
- Other

## 4.6 Rich Content

Rich Content may contain:

- Text
- Images
- External links
- Tables
- Files

Rich Content shall be used wherever educational or informational material requires more than plain text.

---

# 5. Authentication and User Management

## 5.1 Authentication

### FR-AUTH-001 — User Authentication

The system shall allow active users to log in using their registered username and password.

The system shall reject authentication attempts for invalid credentials and accounts that are not permitted to authenticate, including deactivated or suspended accounts.

### FR-AUTH-006 — User Logout

The system shall allow authenticated users to log out.

Logging out shall invalidate the current server-side session so that the same session cannot be used for subsequent authenticated requests.

### FR-AUTH-007 — No Public Registration

The system shall not provide public self-service registration.

All Student, Professor, and Administrator accounts shall be created through an authorized administrative process. The system shall not expose a public account-registration endpoint or workflow.

### FR-AUTH-008 — Administrator-Created Accounts

The system shall allow authorized Administrators to create user accounts for Students and Professors.

Account creation shall be based on information supplied or confirmed by the university's authorized administrative process. The exact operational process for importing or entering university records is outside the scope of this requirement unless separately specified.

### FR-AUTH-009 — Account Activation Invitation

When an Administrator creates an account, the system shall support an invitation or activation process through which the intended user can initialize their account.

The system shall not require Administrators to choose or communicate a user's permanent password. The intended user shall set their own password through the activation process.

An account that has not completed required activation shall not be treated as an Active account and shall not be permitted to use normal authenticated functionality.

### FR-AUTH-010 — Email Verification

The system shall support verification of a user's email address through a verification link or equivalent one-time verification mechanism.

Email verification shall be completed during initial account activation, before the account becomes Active. A verified email address shall be associated with the account only after successful verification.

The system shall not treat an email address as verified merely because an Administrator entered it.

### FR-AUTH-011 — Password Change

An authenticated user shall be able to change their password from account settings.

The user shall provide their current password and a valid new password, except where a separate secure re-authentication mechanism is explicitly defined.

### FR-AUTH-012 — Password Reset Before Login

The system shall provide a password-reset workflow for users who cannot log in because they have forgotten their password.

The workflow shall allow a user to request a reset using their registered account identifier, such as username or email address. The system shall not disclose whether an account exists through distinguishable public responses.

A password reset shall require a valid, unexpired, single-use reset token delivered through the user's verified email address. A successful reset shall invalidate the reset token and existing authenticated sessions for that account.

### FR-AUTH-013 — Password Reset Token Security

Password-reset and account-activation tokens shall be unpredictable, single-use, and time-limited. Expired or previously used tokens shall be rejected.

The exact token lifetime and password-composition policy shall be specified in the security/technical design before implementation.

### FR-AUTH-014 — Email Change and Re-verification

If a user's email address is changed, the new address shall not be treated as verified until the verification process succeeds.

The system shall retain the previous verified status only for the previous address; it shall not transfer that status automatically to the new address.

### FR-AUTH-015 — Server-Side Session Management

The system shall use server-side sessions for browser authentication.

The browser shall carry a session identifier in a cookie configured with appropriate security attributes, including HttpOnly and Secure in production, and an appropriate SameSite policy. Session identifiers shall not be exposed to client-side scripts or placed in URLs.

The server shall validate the session for each authenticated request and enforce authorization independently of client-side UI state.

### FR-AUTH-016 — Session Expiration

The system shall expire sessions according to configured inactivity and absolute-lifetime limits. The specific durations shall be defined in the security/technical design before implementation.

The system shall allow multiple simultaneous sessions for the same account in the MVP unless an Administrator or security policy requires otherwise.

### FR-AUTH-017 — Session Invalidation After Credential or Status Changes

The system shall invalidate existing sessions for an account after a successful password reset and when the account is suspended or deactivated.

After a password change while authenticated, the system shall invalidate other sessions for the account. The current session may remain active only if the implementation securely re-establishes or validates it.

### FR-AUTH-018 — Administrative Account Recovery

If a user loses access to their verified email address or cannot complete automated password reset, the user shall be directed to the university's authorized administrative support process.

Authorized Administrators shall be able to update account contact information and assist with recovery according to administrative permissions. Any new email address shall be verified before it can be used for email-based password reset.

Administrative recovery shall not require a separate automated account-recovery subsystem in the MVP.

### FR-AUTH-019 — Account Status

The system shall distinguish at minimum the following account states:

- Pending Activation
- Active
- Suspended
- Deactivated

Only Active accounts shall be permitted to authenticate and use normal authenticated functionality.

Pending Activation accounts have not completed required initial activation. Suspended accounts are temporarily prevented from authenticating. Deactivated accounts are no longer permitted to authenticate.

### FR-AUTH-020 — Account Deactivation and Historical Data

The system shall allow authorized Administrators to deactivate Student and Professor accounts when the person is no longer eligible to use the platform, including when a Student graduates or a Professor leaves the university.

Deactivation shall prevent future authentication and authenticated actions. Deactivation shall not automatically delete the person's historical academic records, authored reviews, course relationships, or other records that must be retained for platform integrity.

Deactivation shall be distinct from permanent deletion. Any permanent deletion shall be governed by separately defined data-retention and privacy rules.

## 5.2 Roles and Access

### FR-AUTH-002 — User Roles

The system shall support:

- Student
- Teaching Assistant capability
- Professor
- Administrator

A TA shall remain a Student at the global user-role level and receive additional permissions through course assignment.

### FR-AUTH-003 — Role-Based Access

The system shall restrict functionality according to the user's role, account status, and course relationship.

### FR-AUTH-021 — Administrator Privileges

Only authorized Administrators shall be able to create, suspend, reactivate, or deactivate user accounts and manage account-level administrative information.

Administrative account-management actions shall be restricted to authorized administrative users.

## 5.3 Student Profile

### FR-AUTH-004 — Student Profile

The system shall maintain a Student profile containing at minimum:

- Full name
- Profile picture
- Student Code / کد دانشجویی
- Major
- Degree
- University and Department

Additional academic information may be added later.

## 5.4 Professor Profile

### FR-AUTH-005 — Professor Profile

The system shall maintain a Professor profile containing:

- Full name
- Profile picture
- Academic degree / rank
- Field of study / specialization
- Description / biography
- Professional email address(es)
- Relevant social / professional links
- Resume / CV
- Where to teach

## 5.5 Password and Credential Policy

### FR-AUTH-022 — Password Storage

The system shall never store passwords in plaintext. Passwords shall be stored using a suitable, modern password-hashing algorithm and configuration selected during technical design.

### FR-AUTH-023 — Password Validation

The system shall enforce a password policy that discourages easily guessed passwords and rejects passwords that fail the configured minimum security requirements.

The exact length, blocklist, and other password rules shall be specified in the security/technical design before implementation. The system should favor long, unique passwords and avoid relying solely on arbitrary complexity rules.

### FR-AUTH-024 — Authentication Abuse Protection

The system shall apply reasonable protections against repeated automated authentication and password-reset attempts, such as rate limiting or temporary throttling. The exact thresholds and implementation shall be specified in the security/technical design.

# 6. Academic Term Management

## 6.1 Create Academic Terms

### FR-TERM-001 — Academic Term Creation

The system shall support creation of Academic Terms.

Academic Terms may be created:

- Automatically according to the configured academic calendar.
- Manually by an Administrator when necessary.

Each Academic Term shall have a unique Academic Term Code.

### FR-TERM-002 — Automatic Academic Term Creation

The system should be able to automatically create upcoming Academic Terms according to the configured academic calendar.

Creating an Academic Term does not automatically make it the Current Academic Term.

## 6.2 Term Status

The system shall conceptually distinguish:

- Current
- Previous / Historical
- Upcoming

### FR-TERM-003 — View Academic Terms

The Administrator shall be able to view all registered Academic Terms, including:

- Term name
- Academic year
- Academic Term Code
- Status

### FR-TERM-007 — Current Academic Term

The system shall maintain exactly one Current Academic Term.

The Current Academic Term shall be used to determine:

- Current Course Offerings
- Historical Course Offerings
- Students' Current Courses
- Professors' Current Teaching
- Current academic content

### FR-TERM-008 — Historical Terms

Academic Terms earlier than the Current Academic Term according to the configured academic calendar shall be treated as Previous / Historical Terms.

Administrators shall not be required to manually mark them as historical.

### FR-TERM-009 — Upcoming Terms

Created Academic Terms that have not yet become Current shall be treated as Upcoming Terms.

Upcoming Terms may be used to prepare future Course Offerings.

## 6.3 Term Transition

### FR-TERM-005 — Automatic Current Term Transition

The system shall automatically transition the Current Academic Term according to the configured academic calendar.

### FR-TERM-006 — Manual Current Term Override

An Administrator shall be able to manually change the Current Academic Term for exceptional situations such as:

- Academic calendar changes
- Delayed semester starts
- University schedule changes
- Incorrect automatic transition
- Testing
- Exceptional circumstances

Only one Academic Term shall be Current at any time.

## 6.4 Term Updates and Historical Data

### FR-TERM-004 — Update Academic Term

An Administrator shall be able to update Academic Term information when permitted.

Updating a term shall not unintentionally modify associated Course Offerings or historical academic data.

### FR-TERM-010 — Historical Data Preservation

Changing the Current Academic Term shall not modify, delete, or reassign:

- Historical Course Offerings
- Resources
- Reviews
- Student course records
- Other historical academic data

Historical Course Offerings shall remain associated with their original Academic Term.

### FR-TERM-011 — Course Offering Classification

The status of a Course Offering shall be determined from its associated Academic Term.

The system shall not require a manually maintained `is_current` or `is_historical` field for every Course Offering.

---

# 7. Course Management

# 7.1 Course Overview

The platform shall provide a dedicated Course page.

The Course interface shall provide:

- Overview
- Announcements
- Lecture Slides & Materials
- Assignments
- Exams
- Calendar

The page shall display the academic context of the Course, including its Academic Term, instructor, TAs, and other relevant information.

# 7.2 Course Information

Each Course shall maintain:

| Field | Description |
|---|---|
| Course Code | Official university course code |
| Course Name | Official course name |
| Degree Level | Academic degree level |
| Major | Associated academic major / program |
| Department | Associated academic department |
| Year of Entry | Student entry year / cohort |
| Type | Course classification such as required / optional / specialized |
| Prerequisites | One or more prerequisite courses |
| Description | General course description |
| Presentation Time | Official university-assigned time |

## 7.3 Student Course Access

### FR-COURSE-STUDENT-001 — My Courses

A Student shall have access to My Courses.

My Courses shall contain Courses the Student has added and organize them by Academic Term.

### FR-COURSE-STUDENT-002 — Search Courses

A Student shall be able to search for Courses by at least:

- Course Name
- Course Code

Additional search criteria may be added in future versions.

### FR-COURSE-STUDENT-003 — View Search Results

Search results shall provide enough information to identify a Course, including at minimum:

- Course Code
- Course Name

Where applicable, results may also display:

- Degree Level
- Major
- Department
- Academic Term
- Instructor

### FR-COURSE-STUDENT-004 — View Course

A Student shall be able to select a Course from search results and access its Course page according to availability and permissions.

### FR-COURSE-STUDENT-005 — Add Course to My Courses

A Student shall be able to add a Course to My Courses.

Adding a Course shall associate the Student with the Course for access through My Courses and shall not modify the Course itself.

### FR-COURSE-STUDENT-006 — Remove Course from My Courses

A Student shall be able to remove a Course from My Courses.

Removing it shall not delete or modify the Course.

### FR-COURSE-STUDENT-007 — Browse My Courses by Academic Term

A Student shall be able to browse My Courses by:

- Current Term
- Previous Terms

The system shall determine classification from Academic Term configuration.

### FR-COURSE-STUDENT-008 — Access Course Sections

A Student shall be able to access available Course sections:

- Overview
- Announcements
- Lecture Slides & Materials
- Assignments
- Exams
- Calendar

## 7.4 Course Overview — Student

### FR-COURSE-STUDENT-009 — View General Course Information

The Overview shall display:

- Course Code
- Course Name
- Academic Term
- Degree Level
- Major
- Department
- Year of Entry
- Course Type
- Prerequisites
- Presentation Time

### FR-COURSE-STUDENT-010 — View Course Offering Information

The Student shall be able to view:

- Instructor
- TAs
- Meeting Times
- Syllabus
- Objectives
- References / Textbooks
- Grading Policy
- Course Description

The system shall clearly distinguish Presentation Time from actual Meeting Times.

### FR-COURSE-STUDENT-011 — Meeting Times

Students shall be able to view current Meeting Times.

Meeting Times may include:

- Day
- Start Time
- End Time
- Location
- Meeting Type

---

# 8. Course Content and Resources

## 8.1 Announcements

### FR-COURSE-STUDENT-012 — View Announcements

Students shall be able to view published announcements containing:

- Title
- Rich Content
- Author
- Author Type / Role
- Publication Date
- Last Updated Date, where applicable

Professor and TA authors shall be distinguishable.

### FR-COURSE-STUDENT-013 — Announcement Deadlines and Events

An Announcement may optionally contain:

- Deadline
- Event Date
- Event Time
- Event Location

Relevant dates and times shall be represented in the Course Calendar.

## 8.2 Lecture Slides and Materials

### FR-COURSE-STUDENT-014 — View Course Materials

Students shall be able to access published Lecture Slides and Materials.

Materials may contain:

- Text
- Images
- External Links
- Tables
- Files

### FR-COURSE-STUDENT-015 — Download Course Materials

Students shall be able to download files when downloading is permitted.

Content that is not downloadable shall not expose a download option.

## 8.3 Assignments

### FR-COURSE-STUDENT-016 — View Assignments

Students shall be able to view published assignments containing at minimum:

- Title
- Rich Content / Instructions
- Publication Date
- Deadline, where applicable
- Required / Optional status

### FR-COURSE-STUDENT-017 — Download Assignment Files

Students shall be able to download assignment files when available and permitted.

### FR-COURSE-STUDENT-018 — Assignment Calendar Integration

Assignment deadlines shall appear in the Course Calendar and be identified as Assignment deadlines.

## 8.4 Exams

### FR-COURSE-STUDENT-019 — View Exams

Students shall be able to view Exams containing, where applicable:

- Exam title
- Exam date
- Exam time
- Description
- Covered topics / syllabus
- Maximum score
- Rich Content
- Associated files

### FR-COURSE-STUDENT-020 — View Historical Exams

Students shall be able to access historical examinations that have been published by authorized users.

### FR-COURSE-STUDENT-021 — Exam Calendar Integration

Exam dates and times shall appear in the Course Calendar.

---

# 9. Course Calendar

## 9.1 Calendar Scope

### FR-COURSE-STUDENT-022 — View Course Calendar

Students shall be able to access a Course Calendar aggregating:

- Regular Meeting Times
- Make-up Classes
- Assignment Deadlines
- Exams
- Announcements containing scheduled events
- Other Course Events

### FR-COURSE-STUDENT-023 — Meeting Time Calendar Integration

The Calendar shall use Meeting Times when available.

If structured Meeting Times are unavailable, the system may use the official Presentation Time as a fallback schedule.

### FR-COURSE-STUDENT-024 — Calendar Event Information

Each Calendar event shall provide information relevant to its type.

### FR-COURSE-STUDENT-025 — Calendar Event Navigation

Students shall be able to select a Calendar event and navigate to the associated Course content.

Example:

```text
Calendar
  ↓
Assignment 2 — Deadline
  ↓
Assignment 2
```

## 9.2 Calendar as a Derived View

Calendar events may originate from:

```text
Course Calendar
├── Meeting Times
│   ├── Regular Class
│   ├── Make-up Class
│   ├── Office Hour
│   └── Q&A Session
├── Assignments
│   └── Deadline
├── Exams
│   └── Exam Date/Time
├── Announcements
│   └── Scheduled Event / Deadline
└── Other Course Events
```

When an underlying item is created, updated, or deleted, its corresponding Calendar representation shall be updated accordingly.

The system shall avoid requiring users to enter the same event independently in multiple sections.

---

# 10. Course Management by Professor

### FR-COURSE-PROF-001 — Access Teaching Courses

Professors shall be able to access Courses they are assigned to teach and distinguish:

- Current teaching
- Previous teaching

### FR-COURSE-PROF-002 — View Course

Professors shall use the same Course interface available to Students.

### FR-COURSE-PROF-003 — Manage Course Information

Professors shall be able to manage permitted information including:

- Course Description
- Syllabus
- Objectives
- Textbook / References
- Grading Policy
- Meeting Times
- Calendar
- Lecture Slides
- Materials
- Assignments
- Exams
- Announcements

### FR-COURSE-PROF-004 — Restricted Course Information

Professors shall not modify:

- Course Code
- Course Name
- Degree Level
- Major
- Department
- Year of Entry
- Course Type
- Prerequisites
- Academic Term
- Official Presentation Time

These shall be managed by Administrators.

### FR-COURSE-PROF-005 — Manage Meeting Times

Professors shall be able to create, update, or remove actual Meeting Times for:

- Regular classes
- Make-up classes
- Additional sessions
- Office hours
- Q&A sessions
- Other course-related sessions

Changes shall be reflected in the Course Calendar.

Professors shall not modify official Presentation Time.

### FR-COURSE-PROF-006 — Manage Announcements

Professors shall be able to:

- Create announcements
- Edit their announcements
- Delete their announcements
- Associate dates / times
- Add Rich Content
- Associate announcements with Calendar events

The system shall record the Professor as the author.

### FR-COURSE-PROF-007 — Manage Lecture Slides and Materials

Professors shall be able to:

- Add Lecture Slides
- Add Materials
- Edit materials
- Delete materials
- Attach files
- Add external links
- Add Rich Content

### FR-COURSE-PROF-008 — Manage Assignments

Professors shall be able to:

- Create assignments
- Edit assignments
- Delete assignments
- Define deadlines
- Mark assignments as required / optional
- Add Rich Content
- Attach files or links

Assignment deadlines shall automatically appear in the Course Calendar.

### FR-COURSE-PROF-009 — Manage Exams

Professors shall be able to:

- Create Exam records
- Define title
- Define date and time
- Add descriptions
- Specify covered topics
- Specify maximum score where applicable
- Edit exam information
- Delete exam information
- Upload / publish historical exam content

Exam dates shall automatically appear in the Course Calendar.

### FR-COURSE-PROF-010 — Manage Calendar

Professors shall be able to manage permitted Course Calendar events.

Events may originate from:

- Meeting Times
- Assignments
- Exams
- Announcements
- Manually created Course Events

The system shall avoid duplicate manual entry of derived events.

---

# 11. Course Management by TA

A TA is a Student with additional permissions for assigned Courses.

### FR-COURSE-TA-001 — Access Assigned Courses

A TA shall access Courses for which they are assigned and retain normal Student capabilities.

### FR-COURSE-TA-002 — View Course

A TA shall use the same Course interface available to Students.

### FR-COURSE-TA-003 — Manage Course Resources

A TA shall be able to manage permitted:

- Lecture Slides
- Materials
- Assignments
- Announcements
- Course Calendar events

### FR-COURSE-TA-004 — Manage Assignments

A TA may:

- Create assignments
- Edit assignments
- Delete assignments
- Define deadlines
- Mark required / optional
- Add Rich Content
- Attach files or links

unless restricted by the Instructor or Administrator.

### FR-COURSE-TA-005 — Manage Announcements

A TA may:

- Create announcements
- Edit authorized announcements
- Delete authorized announcements
- Add Rich Content
- Associate dates and times

The system shall record the specific TA responsible for creation or modification.

### FR-COURSE-TA-006 — Manage Materials

A TA may:

- Upload Lecture Slides
- Upload Course Materials
- Edit materials
- Delete materials
- Add external links
- Add Rich Content

### FR-COURSE-TA-007 — Manage Meeting Times

A TA may be authorized to manage:

- Q&A sessions
- Office hours
- Make-up classes

A TA shall not modify official Presentation Time.

Regular-class Meeting Time modification shall depend on assigned permissions.

### FR-COURSE-TA-008 — Restricted Information

A TA shall not modify:

- Course Code
- Course Name
- Degree Level
- Major
- Department
- Year of Entry
- Course Type
- Prerequisites
- Academic Term
- Official Presentation Time
- Instructor

A TA shall not assign or remove other TAs unless explicitly authorized.

---

# 12. Course Management by Administrator

### FR-COURSE-ADMIN-001 — Manage Courses

Administrators shall be able to:

- Create Courses
- View Courses
- Update Courses
- Deactivate Courses
- Delete Courses

according to system rules.

### FR-COURSE-ADMIN-002 — Manage Course Information

Administrators shall have full control over permanent Course information:

- Course Code
- Course Name
- Degree Level
- Major
- Department
- Year of Entry
- Type
- Prerequisites
- Description

### FR-COURSE-ADMIN-003 — Manage Course Offerings

Administrators shall manage:

- Academic Term
- Instructor
- TAs
- Official Presentation Time

### FR-COURSE-ADMIN-004 — Manage Course Content

Administrators shall be able to manage Course Offering content and resources when administrative intervention is required, including:

- Syllabus
- Objectives
- Textbooks / References
- Grading Policy
- Meeting Times
- Calendar
- Lecture Slides
- Materials
- Assignments
- Exams
- Announcements

### FR-COURSE-ADMIN-005 — Manage Course Staff

Administrators shall be able to:

- Assign Instructor
- Change Instructor
- Assign TAs
- Remove TAs

Only authorized users may be assigned as Instructor or TA.

### FR-COURSE-ADMIN-006 — Manage Official Presentation Time

Administrators shall be able to create or modify official Presentation Time.

Professors and TAs shall not modify it.

---

# 13. Course Access and Permission Model

### FR-COURSE-ACCESS-001 — Permission Enforcement

The system shall enforce Course permissions according to the user's role and relationship with the Course.

| Capability | Student | TA | Professor | Admin |
|---|---:|---:|---:|---:|
| View Course | ✓ | ✓ | ✓ | ✓ |
| Add to My Courses | ✓ | ✓ | — | — |
| Manage Course Information | ✗ | ✗ | Limited | ✓ |
| Manage Course Resources | ✗ | ✓ | ✓ | ✓ |
| Manage Assignments | ✗ | ✓ | ✓ | ✓ |
| Manage Announcements | ✗ | ✓ | ✓ | ✓ |
| Manage Exams | ✗ | Limited | ✓ | ✓ |
| Manage Calendar | View | Limited | ✓ | ✓ |
| Change Instructor | ✗ | ✗ | ✗ | ✓ |
| Change TAs | ✗ | ✗ | ✗ | ✓ |
| Change Academic Term | ✗ | ✗ | ✗ | ✓ |
| Change Presentation Time | ✗ | ✗ | ✗ | ✓ |
| Change Course Code | ✗ | ✗ | ✗ | ✓ |

The exact TA permission set may be refined before implementation.

---

# 14. Professor Discovery — Ostad Shenasi

## 14.1 Directory Access

### FR-PRF-001 — Access Professor Directory

Authenticated Students shall be able to access Ostad Shenasi.

### FR-PRF-002 — Search Professors

Students shall be able to search Professors by at least:

- Professor full name

Additional search criteria may be added later.

### FR-PRF-003 — Professor Search Results

Results shall provide at minimum:

- Full name
- Profile picture
- Academic rank, where available

### FR-PRF-004 — View Professor Profile

Students shall be able to select a Professor and access the Professor profile.

---

# 15. Professor Profiles

### FR-PRF-005 — Professor Information

A Professor profile shall display:

- Full name
- Profile picture
- Academic rank / degree
- Biography / description
- Professional email address(es)
- Professional / social links
- Resume / CV

Supported external links may include:

- LinkedIn
- Google Scholar
- Telegram
- WhatsApp
- Other relevant professional or communication links

### FR-PRF-006 — Professor Resume

Students shall be able to view or download the Resume / CV when available and permitted.

### FR-PRF-007 — Professor Rating Summary

A Professor profile shall display an overall rating:

- On a five-point scale
- Calculated from published reviews
- Updated when the set of published reviews changes

### FR-PRF-008 — Review Count

A Professor profile shall display the number of published reviews included in the Professor's evaluation.

### FR-PRF-009 — Professor Evaluation Criteria

The profile shall display aggregated scores for evaluation criteria defined by the Administrator.

### FR-PRF-010 — Professor Courses

The profile shall display Courses taught by the Professor and distinguish:

- Current Academic Term
- Previous Academic Terms

Historical information shall be displayed when available.

---

# 16. Professor Reviews — Public View

### FR-REVIEW-001 — View Published Reviews

Students shall be able to view Published reviews associated with a Professor.

Only Published reviews shall be publicly visible.

### FR-REVIEW-002 — Review Information

Each Published review shall display:

- Course
- Academic Term
- Overall rating
- Evaluation criteria and ratings
- Written experience / comment
- Publication date

### FR-REVIEW-003 — Anonymous Reviewer Identity

The identity of the Student who submitted a review shall not be publicly displayed.

The public profile shall not expose:

- Student full name
- Profile picture
- Email
- Student ID
- Other personally identifiable information

Reviewer identity may be retained internally for authentication, moderation, eligibility, and abuse prevention.

### FR-REVIEW-004 — Review Course Context

Each review shall be associated with exactly:

- One Professor
- One Course
- One Academic Term

The Course and Academic Term shall be displayed with the published review.

### FR-REVIEW-005 — Review Date

The system shall record review submission date and time.

The published review shall display the relevant publication date.

The original submission date may also be retained internally.

---

# 17. Professor Review Submission

## 17.1 Submission

### FR-REVIEW-006 — Submit Review

An authenticated Student shall be able to submit a Professor review from the Professor's profile.

### FR-REVIEW-007 — Select Course

The Student shall select the Course associated with their experience.

Only Courses valid for the selected Professor may be selected.

### FR-REVIEW-008 — Select Academic Term

The Student shall select the Academic Term in which they took the selected Course with the Professor.

### FR-REVIEW-009 — Overall Rating

The Student shall provide an overall rating using a five-point scale.

### FR-REVIEW-010 — Evaluation Criteria

The Student shall rate all required evaluation criteria defined by the Administrator.

Each criterion shall have a defined rating scale.

### FR-REVIEW-011 — Professor Features

The Student shall be able to select applicable Professor Features defined globally by the Administrator.

Examples:

- Strict Professor
- Difficult Grading
- Heavy Workload
- Attendance Strictness

### FR-REVIEW-012 — Written Review

The Student shall be able to submit a written description of their experience.

Appropriate validation rules shall apply.

## 17.2 Validation

### FR-REVIEW-013 — Review Submission Validation

Before accepting a review, the system shall validate:

- Student authentication
- Professor selection
- Course selection
- Academic Term selection
- Overall rating
- Required evaluation criteria
- Required fields
- Review eligibility
- Duplicate-review prevention

The system shall prevent duplicates for the same:

```text
Student + Professor + Course + Academic Term
```

### FR-REVIEW-014 — Review Submission Status

A successfully submitted review shall initially receive:

```text
Pending Moderation
```

Pending reviews shall not be publicly visible.

### FR-REVIEW-015 — Review Editing

The system shall define whether Students may edit reviews after submission.

If editing is permitted, the system shall define whether edited reviews return to moderation.

---

# 18. Review Eligibility and Duplicate Prevention

### FR-REVIEW-016 — Review Eligibility

The system shall determine whether a Student is eligible to review a Professor, Course, and Academic Term combination.

The exact verification mechanism is not yet defined.

Possible mechanisms include:

- University enrollment data
- Student confirmation
- Administrator verification
- Other institutional verification

### FR-REVIEW-017 — Duplicate Review Prevention

The system shall prevent more than one review for:

```text
Student + Professor + Course + Academic Term
```

A Student may review the same Professor for different Courses or Academic Terms.

### FR-REVIEW-018 — Review Eligibility Across Courses

A review for one Course shall not prevent reviewing the same Professor for another Course or Academic Term.

---

# 19. Review Moderation

## 19.1 Moderation Workflow

### FR-MOD-001 — Review Moderation Queue

Administrators shall have a moderation interface containing reviews awaiting moderation.

### FR-MOD-002 — Review Details

Administrators shall be able to view complete review content and relevant internal metadata, including reviewer identity where necessary.

Reviewer identity shall remain hidden from normal Students.

### FR-MOD-003 — Publish Review

Administrators shall be able to approve a Pending Moderation review.

Approval changes the review status to:

```text
Published
```

Published reviews become visible to Students.

### FR-MOD-004 — Reject Review

Administrators shall be able to reject reviews.

Rejected reviews shall not be publicly visible.

The system should retain the review and moderation status unless permanent deletion is explicitly required.

### FR-MOD-005 — Hide Published Review

Administrators shall be able to hide previously published reviews.

Hidden reviews shall no longer be visible to Students.

### FR-MOD-006 — Delete Review

Administrators shall be able to permanently delete reviews according to system rules.

Deleted reviews shall no longer contribute to:

- Overall Professor rating
- Evaluation criterion scores
- Published review count

### FR-MOD-007 — Moderation Criteria

Moderation shall address inappropriate content such as:

- Harassment
- Personal attacks
- Spam
- Irrelevant content
- Personal information
- Offensive / inappropriate content
- Other platform-rule violations

Legitimate negative criticism shall not be removed solely because it is negative.

### FR-MOD-008 — Review Status

The system shall distinguish at minimum:

- Pending Moderation
- Published
- Rejected
- Hidden

---

# 20. Review Reporting

### FR-MOD-009 — Report Review

Students shall be able to report a Published review they believe violates moderation rules.

### FR-MOD-010 — Report Reason

Reports shall support at minimum:

- Harassment / Personal Attack
- Spam
- Irrelevant Content
- Personal Information
- Inappropriate Content
- Other

### FR-MOD-011 — Review Report Management

Administrators shall be able to view and manage reports and determine appropriate moderation action.

---

# 21. Professor Administration

### FR-PRF-ADMIN-001 — Create Professor

Administrators shall be able to create Professor profiles and provide available personal and professional information.

### FR-PRF-ADMIN-002 — Update Professor

Administrators shall be able to update:

- Full name
- Profile picture
- Academic rank / degree
- Biography
- Email addresses
- Social / professional links
- Resume / CV

### FR-PRF-ADMIN-003 — Deactivate Professor

Administrators shall be able to deactivate Professor profiles and prevent the associated user account from authenticating when the Professor is no longer eligible to use the platform, including when the Professor leaves the university.

A deactivated Professor shall not appear as active in normal Professor discovery and search. The Professor's historical profile and teaching history may remain available where appropriate, clearly distinguished from active status.

Historical Courses, Course Offerings, Reviews, and other associated academic records shall remain preserved unless removed under a separately defined data-retention and privacy policy.

Deactivating a Professor profile and deactivating the associated login account shall be coordinated so that the account cannot continue to authenticate as an active Professor.

### FR-PRF-ADMIN-004 — Professor Search

Administrators shall be able to search and locate Professor profiles.

### FR-PRF-ADMIN-005 — Professor Management

Administrators shall be able to view and manage Professor profiles through an administrative interface.

---

# 22. Evaluation Criteria and Professor Features

## 22.1 Evaluation Criteria

### FR-CRITERIA-001 — Manage Evaluation Criteria

Administrators shall be able to:

- Create criteria
- Update criteria
- Deactivate criteria
- Manage criteria

Examples:

- Teaching Quality
- Difficulty
- Fairness
- Availability
- Grading

### FR-CRITERIA-002 — Criteria Rating Scale

Administrators shall be able to define the rating scale for evaluation criteria.

## 22.2 Professor Features

### FR-CRITERIA-003 — Manage Professor Features

Professor Features are global platform-level features, not Professor-specific definitions.

Administrators shall be able to:

- Create features
- Update features
- Deactivate features
- Manage features

Examples:

- Strict Professor
- Difficult Grading
- Heavy Workload
- Attendance Strictness

### FR-CRITERIA-004 — Feature Availability

Only active Professor Features shall be selectable in new reviews.

Deactivating a Feature shall not remove it from historical Published reviews.

---

# 23. User Administration

### FR-USER-ADMIN-001 — Create Student

Administrators shall be able to create Student accounts.

### FR-USER-ADMIN-002 — View Student

Administrators shall be able to view Student account information according to administrative permissions.

### FR-USER-ADMIN-003 — Update Student

Administrators shall be able to update Student account information.

### FR-USER-ADMIN-004 — Deactivate Student

Administrators shall be able to deactivate Student accounts, including when a Student graduates or is otherwise no longer eligible to use the platform.

A deactivated Student shall not be able to authenticate or use authenticated Student functionality.

Academic records associated with the Student, including review authorship retained internally for eligibility, moderation, and abuse prevention, shall be preserved according to the platform's data-retention and privacy rules. Public reviews shall remain anonymous as specified in Section 16 and Section 25.

### FR-USER-ADMIN-005 — Manage Account Status

Authorized Administrators shall be able to view account status and perform permitted status transitions, including suspending, deactivating, and reactivating accounts.

Reactivation shall not bypass required activation or email-verification requirements. The system shall record administrative account-status changes for accountability.

### FR-USER-ADMIN-006 — Create Professor Account

Administrators shall be able to create a Professor user account and associated Professor profile through the administrator-controlled account-creation process. The account shall complete required activation and email verification before becoming Active.

### FR-USER-ADMIN-007 — Create Administrator Account

Administrator accounts shall be created only through an authorized, non-public administrative process. Public registration shall not allow a user to obtain the Administrator role.

---

# 24. Professor Rating and Aggregation

### FR-RATING-001 — Calculate Overall Rating

The system shall calculate Professor overall ratings from eligible Published reviews.

The following shall not contribute:

- Pending Moderation reviews
- Rejected reviews
- Hidden reviews

### FR-RATING-002 — Update Overall Rating

The system shall automatically recalculate the overall rating when a Published review is:

- Added
- Removed
- Hidden
- Restored

### FR-RATING-003 — Calculate Criteria Scores

The system shall calculate aggregated scores for active evaluation criteria using applicable Published reviews.

### FR-RATING-004 — Review Count

The system shall calculate the number of Published reviews included in Professor evaluation.

---

# 25. Privacy and Historical Data

### FR-PRIVACY-001 — Anonymous Public Reviews

Student review authors shall remain anonymous to other Students and public users.

### FR-PRIVACY-002 — Internal Reviewer Identification

The system shall retain sufficient information internally to identify review authors for:

- Duplicate prevention
- Review eligibility
- Moderation
- Abuse prevention
- Administrative investigation

### FR-PRIVACY-003 — Historical Review Preservation

Changing Professor profile information shall not unintentionally modify historical review information.

Published reviews shall retain their original Course and Academic Term context.

---

# 26. Review Lifecycle

A review follows this lifecycle:

```text
Student
   │
   ▼
Submit Review
   │
   ▼
Pending Moderation
   │
   ├──────────────► Rejected
   │
   ▼
Published
   │
   ├──────────────► Hidden
   │
   └──────────────► Deleted
```

Only Published reviews contribute to:

- Overall Professor Rating
- Evaluation Criteria Scores
- Published Review Count

The Student's identity remains private throughout the public review lifecycle.

---

# 27. Cross-Cutting Functional Rules

## 27.1 Historical Data Preservation

Changes to current academic state shall not rewrite historical academic records.

This applies particularly to:

- Academic Terms
- Course Offerings
- Course resources
- Reviews
- Student course records
- Professor teaching history

## 27.2 Derived State

Where a status can be derived from another domain object, the system should avoid requiring redundant manually maintained state.

Example:

```text
Course Offering
      │
      ▼
Academic Term
      │
      ▼
Term Status
(Current / Previous / Upcoming)
```

## 27.3 Account Status Enforcement

Account status shall be enforced by the backend, not merely by hiding frontend controls. Pending Activation, Suspended, and Deactivated accounts shall not be able to perform authenticated actions.

Deactivation shall not automatically erase historical records. Where a user account is deactivated, records that must remain for academic history, review integrity, moderation, or duplicate prevention shall retain the necessary internal associations subject to the platform's privacy and retention rules.

## 27.4 Calendar Consistency

Calendar information shall remain synchronized with its source objects.

Examples:

```text
Assignment Deadline
        ↓
Course Calendar
```

```text
Exam Date / Time
        ↓
Course Calendar
```

```text
Meeting Time
        ↓
Course Calendar
```

Changing the source object shall update its corresponding Calendar representation.

## 27.5 Public vs Internal Information

Information required for moderation, eligibility, or abuse prevention may be retained internally even when it must not be exposed publicly.

The strongest example is review author identity.

---

# 28. Functional Permission Summary

| Capability | Student | TA | Professor | Administrator |
|---|---:|---:|---:|---:|
| Authenticate | ✓ | ✓ | ✓ | ✓ |
| View Courses | ✓ | ✓ | ✓ | ✓ |
| Add Course to My Courses | ✓ | ✓ | — | — |
| Manage Course Academic Metadata | ✗ | ✗ | ✗ | ✓ |
| Manage Course Resources | ✗ | ✓ | ✓ | ✓ |
| Manage Assignments | ✗ | ✓ | ✓ | ✓ |
| Manage Announcements | ✗ | ✓ | ✓ | ✓ |
| Manage Exams | ✗ | Limited | ✓ | ✓ |
| Manage Calendar | View | Limited | ✓ | ✓ |
| Manage Academic Terms | ✗ | ✗ | ✗ | ✓ |
| Manage Professors | ✗ | ✗ | ✗ | ✓ |
| Submit Professor Review | ✓ | ✓ | — | — |
| View Published Reviews | ✓ | ✓ | ✓ | ✓ |
| Moderate Reviews | ✗ | ✗ | ✗ | ✓ |
| Report Reviews | ✓ | ✓ | — | — |
| Manage Review Criteria | ✗ | ✗ | ✗ | ✓ |
| Manage Professor Features | ✗ | ✗ | ✗ | ✓ |
| Manage Students | ✗ | ✗ | ✗ | ✓ |

---

# 29. Requirements Traceability

The original requirement identifiers are preserved in this document.

| Domain | Requirement Prefixes |
|---|---|
| Authentication | `FR-AUTH-*` |
| User Account Administration | `FR-USER-ADMIN-*` |
| Academic Terms | `FR-TERM-*` |
| Student Course Access | `FR-COURSE-STUDENT-*` |
| Professor Course Management | `FR-COURSE-PROF-*` |
| TA Course Management | `FR-COURSE-TA-*` |
| Administrator Course Management | `FR-COURSE-ADMIN-*` |
| Course Access / Permissions | `FR-COURSE-ACCESS-*` |
| Professor Discovery | `FR-PRF-*` |
| Professor Administration | `FR-PRF-ADMIN-*` |
| Reviews | `FR-REVIEW-*` |
| Review Moderation | `FR-MOD-*` |
| Evaluation Criteria | `FR-CRITERIA-*` |
| User Administration | `FR-USER-ADMIN-*` |
| Rating Aggregation | `FR-RATING-*` |
| Privacy | `FR-PRIVACY-*` |

---

# 30. Open Requirements / Decisions to Resolve Before Implementation

The current source requirements contain several intentionally unresolved decisions. These should be resolved before implementation rather than silently assumed.

## 30.1 Authentication Mechanism — Resolved for MVP

The authentication approach is defined by Section 5. The following decisions are established for MVP v0.1:

- **Account creation:** All Student, Professor, and Administrator accounts are administrator-created; public self-registration is not supported.
- **Initial setup:** The intended user activates the account and sets their own password rather than receiving a permanent password from an Administrator.
- **Email verification:** Email verification is required during initial activation. Only a verified email address may be used for email-based password reset.
- **Password change:** Authenticated users can change their password from account settings.
- **Password reset:** Users can request a reset before login using their account identifier and a single-use, expiring token delivered to their verified email address. Public responses shall not reveal whether an account exists.
- **Session mechanism:** Browser authentication uses server-side sessions and appropriately secured cookies. Multiple simultaneous sessions are allowed in the MVP.
- **Session invalidation:** Logout invalidates the current session. Password reset, suspension, and deactivation invalidate existing sessions; password changes invalidate other sessions.
- **Account recovery:** Loss of access to the verified email address is handled through authorized university/Administrator support rather than a separate automated recovery subsystem.
- **Account lifecycle:** Accounts may be Pending Activation, Active, Suspended, or Deactivated. Graduating Students and Professors who leave the university may be deactivated.
- **Historical data:** Deactivation prevents authentication but does not automatically delete historical academic records or reviews.

The following implementation parameters remain to be set in the security/technical design before implementation: session inactivity and absolute lifetimes, activation/reset token lifetimes, exact password validation rules, and authentication/password-reset rate limits.

## 30.2 Review Eligibility Verification

`FR-REVIEW-016` explicitly leaves the eligibility mechanism undefined.

A concrete mechanism must be selected before implementation.

## 30.3 Review Editing

`FR-REVIEW-015` leaves post-submission editing behavior undefined.

The team must decide:

- Whether editing is allowed
- Whether edits require moderation again
- Whether previous versions are retained

## 30.4 TA Permission Model

The FRS repeatedly states that exact TA permissions may be refined.

Before implementation, define permissions at the capability level, especially for:

- Exams
- Regular Meeting Times
- Calendar events
- Deleting resources
- Editing resources created by another user

## 30.5 Course vs Course Offering Boundaries

The FRS distinguishes Course and Course Offering conceptually, but implementation-level ownership should be finalized.

In particular, define exactly which fields belong to:

- Course
- Course Offering
- Academic Term

## 30.6 Course Deletion vs Deactivation

Administrator requirements allow both deletion and deactivation, but the preservation rules imply that historical academic data should generally survive.

The exact deletion semantics need to be defined.

## 30.7 Professor Deactivation

The FRS defines deactivation but does not specify all behavior for:

- Existing reviews
- Existing Course Offerings
- Search results
- Historical profiles

The preservation principle should be made explicit in implementation rules.

## 30.8 Calendar Event Model

The FRS defines Calendar as a centralized/derived representation but also permits manually created Course Events.

The implementation should define which events are:

- Derived
- Independently stored
- Editable directly
- Linked to source objects

## 30.9 Rating Calculation

The FRS states that ratings are calculated from Published reviews but does not define the exact aggregation formula.

For example, it does not specify:

- Arithmetic mean vs another aggregation
- Rounding
- Decimal precision
- Minimum review count
- Treatment of missing criterion values

## 30.10 Review Feature Semantics

Professor Features are global and selected per review, but the FRS does not define how they are aggregated or displayed on Professor profiles.

## 30.11 File and Upload Rules

The FRS references CVs, course files, slides, materials, assignment files, and exam files but does not define:

- Maximum file size
- Allowed file types
- Storage rules
- Versioning
- Virus/malware scanning
- File replacement semantics

## 30.12 Search Behavior

Search requirements specify minimum searchable fields but do not define:

- Exact matching vs partial matching
- Case sensitivity
- Persian/English normalization
- Ranking
- Pagination
- Filters
- Sorting

These are implementation decisions unless promoted to functional requirements.

---

# 31. MVP Acceptance Principle

A feature should be considered functionally complete for MVP v0.1 only when:

1. Its functional requirements are implemented.
2. Its authorization rules are enforced.
3. Its relevant data relationships are preserved.
4. Its historical-data behavior is correct.
5. Its derived information remains consistent with its source.
6. Its public/private information boundaries are enforced.
7. Its unresolved requirements have been explicitly decided where implementation depends on them.

---

# 32. Document Status

**Current version:** MVP v0.1

**Primary purpose:** Functional baseline for system design and implementation.

**Next step:** Resolve the Open Requirements / Decisions in Section 30 before freezing the domain model, authorization model, and implementation architecture.
